package routes

import (
	"restaurant-management/server-2/controllers"

	"github.com/gofiber/fiber/v2"
)

func TableRoute(app fiber.Router) {

	router := app.Group("/table")

	router.Get("/", controllers.GetAllTables)
	router.Get("/:id", controllers.GetTable)
	router.Post("/", controllers.CreateTable)
	router.Patch("/:id", controllers.UpdateTable)
	router.Delete("/:id", controllers.DeleteTable)
}
