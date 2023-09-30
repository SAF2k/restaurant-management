package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/saf2k/restaurant-management/server/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.PATCH("/users/:id", controller.UpdateUser())
	incomingRoutes.DELETE("/users/:id", controller.DeleteUser())
	incomingRoutes.POST("/users/logout", controller.Logout())
}
