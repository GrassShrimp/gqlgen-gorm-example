package auth

import (
	"context"
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/grassshrimp/gqlgen-gorm-example/graph/model"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type User struct {
	role model.Role
}

func (user *User) HasRole(role model.Role) bool {
	return user.role == role
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		ctx := context.WithValue(c.Request.Context(), userCtxKey, &User{
			role: (func(_authorization string) model.Role {
				if _authorization == "" || _authorization != "Basic "+base64.StdEncoding.EncodeToString([]byte("admin")) {
					return model.RoleUser
				}

				return model.RoleAdmin
			}(authorization)),
		})

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func ForContext(ctx context.Context) *User {
	return ctx.Value(userCtxKey).(*User)
}
