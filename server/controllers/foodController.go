package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetFoods(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all foods",
	})
}

func GetFood(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get a food",
	})
}

func CreateFood(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create a food",
	})
}

func UpdateFood(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update a food",
	})
}

func DeleteFood(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete a food",
	})
}
