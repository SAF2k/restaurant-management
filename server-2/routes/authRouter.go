package routes

import (
	"restaurant-management/server-2/controllers"
	"restaurant-management/server-2/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes containes all the auth routes
func AuthRoute(app fiber.Router) {

	router := app.Group("/auth")

	router.Post("/signup", controllers.Signup)
	router.Post("/login", controllers.Login)
	router.Get("user", middleware.GetUserAuth, controllers.GetUser)
	router.Get("/users", controllers.GetAllUsers)
}
