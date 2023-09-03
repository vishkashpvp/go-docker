package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vishkash/go-docker/handlers"
	"github.com/vishkash/go-docker/helper"
)

// init function is a special function that is used for initialization purposes
// or performing one-time tasks before the execution of a program or a package.
func init() {
	// Check if the GIN_MODE environment variable is not set to 'gin.ReleaseMode'.
	if os.Getenv("GIN_MODE") != gin.ReleaseMode {
		// If not in release mode, attempt to load environment variables from a .env file.
		if err := godotenv.Load(); err != nil {
			// Log a fatal error message and terminate the program if loading fails.
			log.Fatal("Error loading .env file")
		}
	}
}

// main function is the entry point of the application.
func main() {
	// Create a new Gin router with default middleware.
	r := gin.Default()

	// Define the port on which the server should listen.
	port := 8080

	// Generate the network address in the format "host:port".
	address := helper.NetworkAddress(port)

	// Define and map routes to their respective handlers.
	r.GET("/", handlers.WelcomeMessage)
	r.GET("/ping", handlers.Pong)
	r.GET("/env", handlers.EnvVars)

	// Start the web server and listen on the specified network address.
	if err := r.Run(address); err != nil {
		// Log a fatal error message and terminate the program if server startup fails.
		log.Fatalf("Failed to start server: %v", err)
	}
}
