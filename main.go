package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"server/database"
	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	// Step 2: Basic Server
	router := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	var origins []string
	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
			log.Println("Allowed Origin:", origins[i])
		}
	} else {
		origins = []string{"http://localhost:5173"}
		log.Println("Allowed Origin: http://localhost:5173")
	}

	config := cors.Config{}
	config.AllowOrigins = origins
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	router.Use(gin.Logger())

	var client *mongo.Client = database.Connect()

	if client != nil {
		if err := client.Ping(context.Background(), nil); err != nil {
			log.Fatalf("Failed to reach server: %v", err)
		}
		defer func() {
			err := client.Disconnect(context.Background())
			if err != nil {
				log.Fatalf("Failed to disconnect from MongoDB: %v", err)
			}
		}()
		fmt.Println("Successfully connected and pinged MongoDB.")
	} else {
		log.Println("Database connection is nil. Ensure MONGODB_URI is provided in .env")
	}

	routes.SetupRoutes(router, client)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
