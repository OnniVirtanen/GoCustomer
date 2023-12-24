package main

import (
	"example.com/backend/application/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Setup the API routes
	api.SetupRouter(router)

	// Start the server on a specific port
	router.Run(":3000") // Or use an environment variable or a config file
}
