package utils

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GinContextFromContext returns the Gin context which is stored in native Context provided by stdlib, use it in congunction with GinContextToContext middleware which stores the Gin Context to native Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
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
