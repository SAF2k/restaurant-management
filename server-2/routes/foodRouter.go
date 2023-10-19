package routes

import (
	"restaurant-management/server-2/controllers"
	"restaurant-management/server-2/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func FoodRoutes(app fiber.Router) {

	router := app.Group("/food").Use(middleware.Auth)

	router.Get("/", controllers.GetAllFood)
	router.Get("/:id", controllers.GetFood)
	router.Get("/menu/:id", controllers.GetFoodByMenu)
	router.Post("/", controllers.CreateFood)
	router.Patch("/:id", controllers.UpdateFood)
	router.Delete("/:id", controllers.DeleteFood)
}
