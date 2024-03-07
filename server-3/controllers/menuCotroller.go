package controllers

import (
	"fmt"
	"restaurant-management/server-2/database"
	"restaurant-management/server-2/models"
	"restaurant-management/server-2/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetAllMenus(ctx *fiber.Ctx) error {
	//Collect store id from params
	storeId := ctx.Params("s_id")

	//Find all menus by store id
	result, err := menuCollection.Find(ctx.Context(), bson.M{"store_id": storeId})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	//Create bson.M type var
	var menus []bson.M

	//Decode all menus
	if err := result.All(ctx.Context(), &menus); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any menu")
	}

	//Return menu
	return ctx.JSON(menus)
}

func GetMenu(ctx *fiber.Ctx) error {
	//Collect menu id and store id from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create menu model
	var menu models.Menu

	//Find menu by menu id and store id
	err := menuCollection.FindOne(ctx.Context(), bson.M{"store_id": storeId, "menu_id": menuId}).Decode(&menu)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Menu not found")
	}

	//Return menu
	return ctx.JSON(menu)
}

func CreateMenu(ctx *fiber.Ctx) error {
	//Collect store id from params
	storeId := ctx.Params("s_id")

	//Create a new Menu type
	menu := new(models.Menu)

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, menu); err != nil {
		return err
	}

	menu.ID = primitive.NewObjectID()
	menu.Menu_id = menu.ID.Hex()

	//Create a new menu model
	menuModel := models.Menu{
		ID:         menu.ID,
		Name:       menu.Name,
		Category:   menu.Category,
		Menu_id:    menu.Menu_id,
		Store_id:   &storeId,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	//Insert menu model into database
	result, err := menuCollection.InsertOne(ctx.Context(), menuModel)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create menu")
	}

	//Return menu
	return ctx.JSON(result)
}

func UpdateMenu(ctx *fiber.Ctx) error {
	//Collect menu id and store id from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create a new Menu type
	menu := new(models.Menu)

	//Find menu by menu id and store id
	err := menuCollection.FindOne(ctx.Context(), bson.M{"store_id": storeId, "menu_id": menuId}).Decode(&menu)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Menu not found")
	}

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, menu); err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":       menu.Name,
			"category":   menu.Category,
			"updated_at": time.Now(),
		},
	}

	//Update menu by menu id and store id
	result, err := menuCollection.UpdateOne(ctx.Context(), bson.M{"store_id": storeId, "menu_id": menuId}, update)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update menu")
	}
	fmt.Println(result)

	//Return menu
	return ctx.JSON(result)
}

func DeleteMenu(ctx *fiber.Ctx) error {
	//Collect menu id and store id from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Delete menu by menu id and store id
	result, err := menuCollection.DeleteOne(ctx.Context(), bson.M{"store_id": storeId, "menu_id": menuId})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete menu")
	}

	//Return menu
	return ctx.JSON(result)
}
