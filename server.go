package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/grassshrimp/gqlgen-gorm-example/auth"
	"github.com/grassshrimp/gqlgen-gorm-example/graph"
	"github.com/grassshrimp/gqlgen-gorm-example/graph/generated"
	"github.com/grassshrimp/gqlgen-gorm-example/graph/model"
)

const defaultPort = "8080"

var db *gorm.DB
var cfg *Config

func init() {
	configfile, err := ParseFlags()

	if err != nil {
		fmt.Println(err)
		panic("failed to parse configfile")
	}

	cfg, err = NewConfig(configfile)

	if err != nil {
		fmt.Println(err)
		panic("failed to parse configfile")
	}

	dsn := cfg.MysqlDsn
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()
	// r.Use(dataloader.CustomerLoaderMiddleware(db))
	r.Use(auth.AuthMiddleware())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":" + cfg.Port)
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	cache, err := NewCache(cfg.Redis.Host+":"+strconv.Itoa(cfg.Redis.Port), cfg.Redis.Password, cfg.Redis.Index, 24*time.Hour)

	if err != nil {
		log.Fatalf("cannot create APQ redis cache: %v", err)
	}

	c := generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}

	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		if !auth.ForContext(ctx).HasRole(role) {
			// block calling the next resolver
			return nil, fmt.Errorf("Access denied")
		}

		// or let it pass through
		return next(ctx)
	}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	h.AddTransport(transport.POST{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: cache,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
