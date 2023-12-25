package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"example.com/backend/application/api"
	"example.com/backend/application/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func initDatabase() {
	fmt.Println("Starting sql-driver...")

	// Use an environment variable for the database connection string.
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is not set.")
	}

	// Open up our database connection.
	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Ping the database to check for connection errors
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	fmt.Println("Successfully connected to the database")
}

func initRouter() {
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

	// Use an environment variable for the database connection string.
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("SERVER_PORT environment variable is not set.")
	}
	// Start the server on a specific port
	router.Run(":" + serverPort) // Or use an environment variable or a config file
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initDatabase()
	initRouter()
}
