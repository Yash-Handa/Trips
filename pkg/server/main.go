package server

import (
	"log"

	"github.com/Yash-Handa/Trips/internal/handlers"
	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Server configs
var host, port string

func init() {
	host = utils.MustGet("HOST")
	port = utils.MustGet("PORT")
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
