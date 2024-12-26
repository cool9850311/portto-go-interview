package main

import (
	"Go-Service/src/main/infrastructure/router"
	"github.com/gin-gonic/gin"
	"Go-Service/src/main/infrastructure/initializer"
)

func main() {
	initializer.InitLog()
	initializer.InitConfig(initializer.Log)
	initializer.InitMongoClient()
	// Create a new Gin engine
	r := gin.Default()

	// Setup the router
	router.SetupRouter(r, initializer.DB, initializer.Log)

	// Start the server on port 8080
	r.Run(":8080")
}
