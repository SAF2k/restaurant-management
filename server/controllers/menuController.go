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

		if menu.Start_Date != nil && menu.End_Date != nil {
			if inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Start date and end date are not valid",
				})
				defer cancel()
				return
			}

			updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.Start_Date})
			updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.End_Date})
		}

		if menu.Name != nil {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}

		if menu.Category != nil {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
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

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
