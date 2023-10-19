package main

import (
	"fmt"
	"restaurant-management/server-2/config"
	"restaurant-management/server-2/routes"
	"restaurant-management/server-2/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	allowSites := config.ALLOW_SITES

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowOrigins:     allowSites,
	}))

	routes.AuthRoute(app)
	routes.StoreRoute(app)
	routes.MenuRoute(app)
	routes.FoodRoutes(app)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
