package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupRoutes(router *gin.Engine, client *mongo.Client) {
	// Let's keep the simple hello route here for testing
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Movie API is running!")
	})

	// Public routes
	router.GET("/movies", controllers.GetMovies(client))
	router.GET("/genres", controllers.GetGenres(client))

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleWare())
	{
		protected.GET("/movie/:imdb_id", controllers.GetMovie(client))
		protected.POST("/movie", controllers.AddMovie(client))
		protected.GET("/recommendedmovies", controllers.GetRecommendedMovies(client))
		protected.PATCH("/updatereview/:imdb_id", controllers.AdminReviewUpdate(client))
	}

	// User authentication routes
	router.POST("/register", controllers.RegisterUser(client))
	router.POST("/login", controllers.LoginUser(client))
	router.POST("/logout", controllers.LogoutHandler(client))
	router.POST("/refresh", controllers.RefreshTokenHandler(client))
}
