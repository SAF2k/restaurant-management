package routes

import (
	"restaurant-management/server-2/controllers"
	"restaurant-management/server-2/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func FoodRoutes(app fiber.Router) {

	router := app.Group("/food")

	router.Get("/", controllers.GetAllFood)
	router.Get("/:id", controllers.GetFood)
	router.Get("/menu/:id", controllers.GetFoodByMenu)
	router.Post("/", middleware.Auth, controllers.CreateFood)
	router.Patch("/:id", middleware.Auth, controllers.UpdateFood)
	router.Delete("/:id", middleware.Auth, controllers.DeleteFood)
}
