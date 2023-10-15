package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/saf2k/restaurant-management/server/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := setupRouter()
	router.Run(":" + port)
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(CorsMiddleware())

	registerRoutes(router)

	return router
}

func registerRoutes(router *gin.Engine) {
	routes.UserRoutes(router)

	// router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)

	// Uncomment the following line to enable authentication middleware
}

func CorsMiddleware() gin.HandlerFunc {
	// Create a new CORS middleware with specific options
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Replace with your frontend's origin
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return func(c *gin.Context) {
		// Apply the customized CORS middleware to the context
		corsConfig.HandlerFunc(c.Writer, c.Request)
		c.Next()
	}
}
