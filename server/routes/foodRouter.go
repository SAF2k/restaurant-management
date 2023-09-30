package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods)
	incomingRoutes.GET("/food/:food_id", controllers.GetFood)
	incomingRoutes.POST("/food", controllers.CreateFood)
	incomingRoutes.PATCH("/food/:id", controllers.UpdateFood)
	incomingRoutes.DELETE("/food/:id", controllers.DeleteFood)
}
