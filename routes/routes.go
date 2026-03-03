package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupRoutes(router *gin.Engine, client *mongo.Client) {
	// Let's keep the simple hello route here for testing
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Movie API is running!")
	})

	// Add more routing logic here similar to MagicStream
	// Need to create controllers for these in the future
}
