package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all users",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get a user",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create a user",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update a user",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete a user",
	})
}
