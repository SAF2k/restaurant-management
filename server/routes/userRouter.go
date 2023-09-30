package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.GET("/users/:user_id", controllers.GetUser)
	incomingRoutes.POST("/users", controllers.CreateUser)
	incomingRoutes.PATCH("/users/:id", controllers.UpdateUser)
	incomingRoutes.DELETE("/users/:id", controllers.DeleteUser)
}
