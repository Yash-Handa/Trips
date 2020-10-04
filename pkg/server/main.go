package server

import (
	"log"

	"github.com/Yash-Handa/Trips/internal/db"
	"github.com/Yash-Handa/Trips/internal/handlers"
	"github.com/Yash-Handa/Trips/internal/middlewares"
	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/gin-contrib/cors"
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
	// connect to db
	db.Connect()
	defer db.Close()

	r := gin.Default()

	// middlewares on global router

	// cors setup
	// - config.AllowAllOrigins = true
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	r.Use(cors.Default())
	r.Use(middlewares.GinContextToContext())
	r.Use(middlewares.Auth(db.GetDB()))

	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET("/", handlers.PlaygroundHandler("/graphql"))
	}

	// GraphQL Handler
	gqlH := handlers.GraphqlHandler(db.GetDB())
	r.GET("/graphql", gqlH)
	r.POST("/graphql", gqlH)

	log.Fatalln(r.Run(host + ":" + port))
}
