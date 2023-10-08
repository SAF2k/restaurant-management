package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/saf2k/restaurant-management/server/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/table/:table_id", controller.GetTable())
	incomingRoutes.POST("/table", controller.CreateTable())
	incomingRoutes.PATCH("/table/:table_id", controller.UpdateTable())
	incomingRoutes.DELETE("/table/:table_id", controller.DeleteTable())
}
