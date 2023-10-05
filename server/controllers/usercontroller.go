package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/database"
	"github.com/saf2k/restaurant-management/server/models"
	"github.com/saf2k/restaurant-management/server/validate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := userCollection.Find(ctx, bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Users not found",
			})
			return
		}
		var allUsers []bson.M

		if err = result.All(ctx, &allUsers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Users not found",
			})
			return
		}

		c.JSON(http.StatusOK, allUsers)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id") // Extract the user ID from the request URL

		// Convert the user ID to an ObjectID (assuming MongoDB)
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Query your database to find the user by their ID
		var user models.User
		filter := bson.M{"_id": objectID}
		err = userCollection.FindOne(c.Request.Context(), filter).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update a user",
		})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete a user",
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var user models.User

		// if err := c.ShouldBindJSON(&user); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// // Check if the user exists in the simulated database
		// hashedPassword, ok := users[user.Username]
		// if !ok || !VerifyPassword(hashedPassword, HashPassword(user.Password)) {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		// 	return
		// }

		// // Generate a token (you should use a proper token library)
		// token := GenerateToken(user.Username)

		// c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Logout",
		})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the user struct using the validator
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.ID = primitive.NewObjectID()
		user.Create_at = time.Now()
		user.Update_at = time.Now()

		// Save the user to your database (not implemented in this example)

		c.JSON(http.StatusOK, user)
	}
}

func HashPassword() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	return true, "Verify Password"
}
