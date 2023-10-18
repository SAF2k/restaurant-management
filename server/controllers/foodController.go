package controller

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saf2k/restaurant-management/server/database"
	"github.com/saf2k/restaurant-management/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := foodCollection.Find(ctx, bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Foods not found",
			})
			return
		}
		var allFoods []bson.M

		if err = result.All(ctx, &allFoods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Foods not found",
			})
			return
		}

		fmt.Println(allFoods)

		c.JSON(http.StatusOK, allFoods)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Food not found",
			})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func GetFoodsByMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var foods []models.Food

		result, err := foodCollection.Find(ctx, bson.M{"menu_id": menuId})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Foods not found",
			})
			return
		}

		if err = result.All(ctx, &foods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Foods not found",
			})
			return
		}
		c.JSON(http.StatusOK, foods)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu
		defer cancel()
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		validationError := validate.Struct(food)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError.Error(),
			})
			return
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Menu not found",
			})
			return
		}
		food.Menu_Name = menu.Category
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": insertErr.Error(),
			})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get food ID from URL parameter
		foodID := c.Param("food_id")

		// Define the filter to find the food by ID
		filter := bson.M{"food_id": foodID}

		// Check if the food item with the specified food_id exists
		var existingFood models.Food
		err := foodCollection.FindOne(c.Request.Context(), filter).Decode(&existingFood)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Food not found",
			})
			return
		}

		// Parse JSON request into a food struct
		var food models.Food
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Create an update document with the $set operator
		update := bson.M{
			"$set": bson.M{
				"name":       food.Name,
				"price":      food.Price,
				"food_image": food.Food_Image,
				"updated_at": time.Now(),
			},
		}

		// Define options for the update operation
		opts := options.Update()

		// Perform the update operation
		result, err := foodCollection.UpdateOne(c.Request.Context(), filter, update, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Food updated successfully",
			"result":  result,
		})
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		result, err := foodCollection.DeleteOne(ctx, bson.M{"food_id": foodId})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Food not found",
			})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
