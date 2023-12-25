package main

import (
	"log"
	"os"

	"example.com/backend/application/api"
	"example.com/backend/application/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Open a file for logging
	file, err := os.OpenFile("log/requests.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Create a new logger instance
	logger := log.New(file, "", log.LstdFlags)

	// Register the middleware
	router.Use(middleware.RequestLoggerMiddleware(logger))

	// Setup the API routes
	api.SetupRouter(router)

	// Start the server on a specific port
	router.Run(":3000") // Or use an environment variable or a config file
}
