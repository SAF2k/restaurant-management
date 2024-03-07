package routes

import (
	"restaurant-management/server-2/controllers"

	"github.com/gofiber/fiber/v2"
)

func MenuRoute(app fiber.Router) {

	router := app.Group("/menu")

	router.Get("/", controllers.GetAllMenus)
	router.Get("/:id", controllers.GetMenu)
	router.Get("/:id/foods", controllers.GetFoodByMenu)
	router.Post("/", controllers.CreateMenu)
	router.Patch("/:id", controllers.UpdateMenu)
	router.Delete("/:id", controllers.DeleteMenu)
	// router.Delete("/:id/foods", controllers.DeleteFoodByMenu)
}
