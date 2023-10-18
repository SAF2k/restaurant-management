package main

import (
	"fmt"
	"restaurant-management/server-2/config"
	"restaurant-management/server-2/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	allowSites := config.ALLOW_SITES

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowOrigins:     allowSites,
	}))

	routes.FoodRoutes(app)
	routes.AuthRoutes(app)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
