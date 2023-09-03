package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Responds with a welcome message.
func WelcomeMessage(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the server")
}

// Responds with a "pong" message.
func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// Retrieves and responds with environment variables.
// It returns a JSON object containing application name and version from environment variables.
func EnvVars(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"App":     os.Getenv("APP_NAME"),
		"Version": os.Getenv("APP_VERSION"),
	})
}
