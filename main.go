package main

import (
	"log"

	"github.com/gin-gonic/gin"

	routes "github.com/cavdy-play/go_mongo/routes"
)

func main()  {
	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
