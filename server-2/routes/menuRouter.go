package routes

import (
	"restaurant-management/server-2/controllers"
	"restaurant-management/server-2/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func MenuRoute(app fiber.Router) {

	router := app.Group("/:s_id/menu")

	router.Get("/", controllers.GetAllMenus)
	router.Get("/:id", controllers.GetMenu)
	// router.Get("/:id/foods", controllers.GetFoodByMenu)
	router.Post("/", controllers.CreateMenu).Use(middleware.Auth)
	// router.Post("/:id/foods", controllers.CreateFood)
	// router.Put("/:id", controllers.UpdateMenu)
	// router.Delete("/:id", controllers.DeleteMenu)
	// router.Delete("/:id/foods", controllers.DeleteFoodByMenu)
}
