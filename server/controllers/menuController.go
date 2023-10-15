package controller

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/database"
	"github.com/saf2k/restaurant-management/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

		// Set the "value" field to a lowercase and no-space string of "category"
		if menu.Category != nil {
			value := strings.ToLower(strings.ReplaceAll(*menu.Category, " ", ""))
			menu.Value = &value
		}
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		defer cancel()
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		if err := menuCollection.FindOne(ctx, filter).Decode(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Menu not found",
			})
			return
		}

		var updateObj primitive.D

		if menu.Name != nil {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}

		if menu.Category != nil {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})

			// Set the "Value" field to a lowercase and no-space string of "category"
			value := strings.ToLower(strings.ReplaceAll(*menu.Category, " ", ""))
			updateObj = append(updateObj, bson.E{Key: "value", Value: value})
		}

		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.Updated_at})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := menuCollection.UpdateOne(ctx, filter, bson.D{
			{Key: "$set", Value: updateObj},
		}, &opt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
		defer cancel()
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
