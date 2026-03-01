package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Step 2: Basic Server
	router := gin.Default()

	// Define a simple /hello route
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
