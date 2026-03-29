package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupRoutes(router *gin.Engine, client *mongo.Client) {
	// Let's keep the simple hello route here for testing
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Movie API is running!")
	})

	// Movie routes
	router.GET("/movies", controllers.GetMovies(client))
	router.GET("/movie/:imdb_id", controllers.GetMovie(client))
	router.POST("/movie", controllers.AddMovie(client))
}
