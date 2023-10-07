package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/database"
	"github.com/saf2k/restaurant-management/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
			c.JSON(400, gin.H{
				"message": "Menu not found",
			})
			return
		}
		c.JSON(http.StatusOK, menu)
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		validationError := validate.Struct(menu)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError.Error(),
			})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		fmt.Println(menu)

		result, insertErr := menuCollection.InsertOne(ctx, menu)

		if insertErr != nil {
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":         msg,
				"error_context": insertErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}

func UpdateMenuItem() gin.HandlerFunc {
	// return func(c *gin.Context) {
	// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	// 	var menu models.Menu

	// 	if err := c.BindJSON(&menu); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	menuId := c.Param("menu_id")
	// 	filter := bson.M{"menu_id": menuId}

	// 	var updateObj primitive.D

	// 	if menu.Start_Date != nil && menu.End_Date != nil {
	// 		if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
	// 			c.JSON(http.StatusInternalServerError, gin.H{
	// 				"error": "Start date and end date are not valid",
	// 			})
	// 			return
	// 		}

	// 		updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
	// 		updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

	// 		if menu.Menu_Name != "" {
	// 			updateObj = append(updateObj, bson.E{"menu_name", menu.Menu_Name})
	// 		}

	// 		if menu.Category != "" {
	// 			updateObj = append(updateObj, bson.E{"menu_name", menu.Category})
	// 		}

	// 		update = append(update, bson.E{"$set", updateObj})

	// 		err := menuCollection.FindOneAndUpdate(ctx, filter, update).Decode(&menu)

	// 		if err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{
	// 				"error": "Menu item was not updated",
	// 			})
	// 			return
	// 		}

	// 		c.JSON(http.StatusOK, menu)
	// 		defer cancel()
	// }
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update a menu item",
		})
	}
}

func DeleteMenuItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menu_id := c.Param("menu_id")
		defer cancel()

		filter := bson.M{"menu_id": menu_id}

		result, err := menuCollection.DeleteOne(ctx, filter)

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
