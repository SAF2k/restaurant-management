package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/routes"
)

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	// router.Use(middleware.Authentication())

	routes.FoodRoutes(router)

	router.Run(":" + port)
}
