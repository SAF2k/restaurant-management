package main

import (
	"fmt"
	"restaurant-management/server-2/config"
	"restaurant-management/server-2/routes"
	"restaurant-management/server-2/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	allowSites := config.ALLOW_SITES

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		// AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH",
		AllowOrigins: allowSites,
	}))

	setupRoutes(app)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}

func setupRoutes(app *fiber.App) {
	routes.AuthRoute(app)
	routes.StoreRoute(app)

	api := app.Group("/:s_id")

	routes.MenuRoute(api)
	routes.FoodRoutes(api)
	routes.TableRoute(api)
}
