package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Yash-Handa/Trips/internal/gql/generated"
	"github.com/Yash-Handa/Trips/internal/gql/resolver"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	c := generated.Config{
		Resolvers: &resolver.Resolver{},
	}

	// custom complexity calculator for query fields which return arrays of certain numbers(count)
	countComplexity := func(childComplexity, count int) int {
		return count * childComplexity
	}
	_ = countComplexity
	// c.Complexity.Query.Users = countComplexity

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	// add Fixed complexity to the server
	// h.Use(extension.FixedComplexityLimit(5))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
