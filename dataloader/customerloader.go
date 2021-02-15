package dataloader

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grassshrimp/gqlgen-gorm-example/graph/generated"
	"github.com/grassshrimp/gqlgen-gorm-example/graph/model"
	"gorm.io/gorm"
)

const loadersKey = "dataloaders"

type Loaders struct {
	CustomerById *generated.CustomerLoader
}

func CustomerLoaderMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), loadersKey, &Loaders{
			CustomerById: generated.NewCustomerLoader(generated.CustomerLoaderConfig{
				MaxBatch: 100,
				Wait:     1 * time.Millisecond,
				Fetch: func(ids []int) ([]*model.Customer, []error) {
					customers := []*model.Customer{}
					fmt.Println(ids)
					result := db.Find(&customers, "CustomerNumber IN ?", ids)

					if result.Error != nil {
						panic(result.Error)
					}

					return customers, nil
				},
			}),
		})
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
