package routes

import (
	"github.com/SAF2k/restaurant-management/server/controllers"

	"github.com/gofiber/fiber/v2"
)

func StoreRoute(app fiber.Router) {

	router := app.Group("/store")

	router.Get("/", controllers.GetAllStores)
	router.Get("/:id", controllers.GetStore)
	router.Post("/", controllers.CreateStore)
	router.Patch("/:id", controllers.UpdateStore)
	router.Delete("/:id", controllers.DeleteStore)
}
