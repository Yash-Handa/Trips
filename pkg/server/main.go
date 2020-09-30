package server

import (
	"log"

	"github.com/Yash-Handa/Trips/internal/handlers"
	"github.com/Yash-Handa/Trips/internal/middlewares"
	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Server configs
var host, port string
var isPgEnabled bool

func init() {
	host = utils.MustGet("HOST")
	port = utils.MustGet("PORT")
	isPgEnabled = utils.MustGetBool("GQL_PLAYGROUND_ENABLED")
}

// Run web server
func Run() {
	r := gin.Default()

	// middlewares on global router
	r.Use(middlewares.GinContextToContext())

	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET("/", handlers.PlaygroundHandler("/graphql"))
	}
	r.POST("/graphql", handlers.GraphqlHandler())

	log.Fatalln(r.Run(host + ":" + port))
}
