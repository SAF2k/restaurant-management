package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/database"
	"github.com/saf2k/restaurant-management/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MenuUpdateRequest struct {
	Name     string `json:"name" binding:"min=2,max=100"`
	Category string `json:"category" binding:"min=2,max=100"`
}

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var menu models.Menu

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Menu not found",
			})
			return
		}

		// Create a MenuResponse with only the desired fields
		response := models.MenuResponse{
			Name:     *menu.Name,
			Category: *menu.Category,
		}

		c.JSON(http.StatusOK, response)
	}
}
func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(ctx, bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Menus not found",
			})
			return
		}
		var allMenus []bson.M

		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allMenus)
	}
}

func CreateMenuItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout and defer canceling it
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Parse JSON request into a menu struct
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Validate the menu struct
		validationError := validate.Struct(menu)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError.Error(),
			})
			return
		}

		// Calculate the current time
		now := time.Now()
		menu.Created_at = now
		menu.Updated_at = now

		// Generate a new ObjectID and set the Menu_id
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		// Insert the menu into the database
		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while creating a new menu item",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func UpdateMenuItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get menu ID from URL parameter
		menuID := c.Param("menu_id")

		// Check if menu_id is empty or not provided
		if menuID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "menu_id is required",
			})
			return
		}

		// Parse JSON request into a menu update request struct
		var updateReq MenuUpdateRequest
		if err := c.ShouldBindJSON(&updateReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Define the filter to find the menu by ID
		filter := bson.M{"menu_id": menuID}

		// Check if the menu item with the specified menu_id exists
		var existingMenu models.Menu
		err := menuCollection.FindOne(c.Request.Context(), filter).Decode(&existingMenu)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Menu not found",
			})
			return
		}

		// Set the updated fields
		updateFields := bson.M{
			"name":       updateReq.Name,
			"category":   updateReq.Category,
			"updated_at": time.Now(),
		}

		// Create an update document with the $set operator
		update := bson.M{
			"$set": updateFields,
		}

		// Define options for the update operation
		opts := options.Update()

		// Perform the update operation
		result, err := menuCollection.UpdateOne(c.Request.Context(), filter, update, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Menu updated successfully",
			"result":  result,
		})
	}
}

func DeleteMenuItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menu_id := c.Param("menu_id")
		defer cancel()

		filter := bson.M{"menu_id": menu_id}

		res_menu, err := menuCollection.DeleteOne(ctx, filter)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Menu item was not deleted",
			})
			return
		}

		res_food, err := foodCollection.DeleteMany(ctx, bson.M{"menu_id": menu_id})

		var result = bson.M{
			"menu": res_menu,
			"food": res_food,
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Menu item was not deleted",
			})
			return
		}

		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}
