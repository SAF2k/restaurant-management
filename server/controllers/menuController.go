package controller

import "github.com/gin-gonic/gin"

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all menu",
		})
	}
}

func GetMenuItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get a menu item",
		})
	}
}

func CreateMenuItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create a menu item",
		})
	}
}

func UpdateMenuItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update a menu item",
		})
	}
}

func DeleteMenuItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete a menu item",
		})
	}
}
