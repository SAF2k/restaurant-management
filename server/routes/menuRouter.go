package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/saf2k/restaurant-management/server/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controller.GetMenu())
	incomingRoutes.GET("/menu/:menu_id", controller.GetMenuItem())
	incomingRoutes.POST("/menu", controller.CreateMenuItem())
	incomingRoutes.PATCH("/menu/:menu_id", controller.UpdateMenuItem())
	incomingRoutes.DELETE("/menu/:menu_id", controller.DeleteMenuItem())
}
