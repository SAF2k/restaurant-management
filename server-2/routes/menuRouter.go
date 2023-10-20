package routes

import (
	"restaurant-management/server-2/controllers"
	"restaurant-management/server-2/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func MenuRoute(app fiber.Router) {

	router := app.Group("/menu")

	router.Get("/", controllers.GetAllMenus)
	router.Get("/:id", controllers.GetMenu)
	// router.Get("/:id/foods", controllers.GetFoodByMenu)
	router.Post("/", middleware.Auth, controllers.CreateMenu)
	// router.Post("/:id/foods", controllers.CreateFood)
	router.Patch("/:id", middleware.Auth, controllers.UpdateMenu)
	router.Delete("/:id", middleware.Auth, controllers.DeleteMenu)
	// router.Delete("/:id/foods", controllers.DeleteFoodByMenu)
}
