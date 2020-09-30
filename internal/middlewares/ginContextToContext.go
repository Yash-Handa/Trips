package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
)

// GinContextToContext adds Gin Context to the native Context provided by golang stdlib, which is then used by gqlgen resolvers
func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
