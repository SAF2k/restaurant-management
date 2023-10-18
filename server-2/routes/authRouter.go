package routes

import (
	"restaurant-management/server-2/controllers"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes containes all the auth routes
func AuthRoutes(app fiber.Router) {

	router := app.Group("/auth")

	router.Post("/signup", controllers.Signup)
	router.Post("/login", controllers.Login)
}
