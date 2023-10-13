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
	// Create a CORS middleware with default options
	corsMiddleware := cors.Default()

	return func(c *gin.Context) {
		// Apply CORS middleware to the context
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	}
}
