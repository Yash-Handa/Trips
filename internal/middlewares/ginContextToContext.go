package middlewares

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GinContextKey is the key in context corresponding to gin context
const GinContextKey = "GinContextKey"

// GinContextToContext adds Gin Context to the native Context provided by golang stdlib, which is then used by gqlgen resolvers
func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GinContextFromContext returns the Gin context which is stored in native Context provided by stdlib, use it in congunction with GinContextToContext middleware which stores the Gin Context to native Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey)
	if ginContext == nil {
		return nil, fmt.Errorf("could not retrieve gin.Context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, fmt.Errorf("gin.Context has wrong type")
	}
	return gc, nil
}

//Usage:
//
// func (r *resolver) Todo(ctx context.Context) (*Todo, error) {
// 	gc, err := GinContextFromContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// }
