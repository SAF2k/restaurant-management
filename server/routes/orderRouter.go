package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/saf2k/restaurant-management/server/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/order/:order_id", controller.GetOrder())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PATCH("order/:order_id", controller.UpdateOrder())
	incomingRoutes.DELETE("/order/:order_id", controller.DeleteOrder())
}
