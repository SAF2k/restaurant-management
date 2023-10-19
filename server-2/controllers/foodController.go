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

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetAllFood(ctx *fiber.Ctx) error {
	//Collect store id from params
	// storeId := ctx.Params("s_id")

	//Find all foods by store id
	result, err := foodCollection.Find(ctx.Context(), bson.M{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	//Create food model
	var foods []models.Food

	//Decode all foods
	if err := result.All(ctx.Context(), &foods); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any food")
	}

	//Return food
	return ctx.JSON(foods)
}

func GetFood(ctx *fiber.Ctx) error {
	//Collect food id and store id from params
	foodId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create food model
	var food models.Food

	//Find food by food id and store id
	err := foodCollection.FindOne(ctx.Context(), bson.M{"store_id": storeId, "food_id": foodId}).Decode(&food)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Food not found")
	}

	//Return food
	return ctx.JSON(food)
}

func GetFoodByMenu(ctx *fiber.Ctx) error {
	//Collect menu id and store id from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Find all foods by menu id and store id
	result, err := foodCollection.Find(ctx.Context(), bson.M{"menu_id": menuId, "store_id": storeId})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	//Create food model
	var foods []models.Food

	//Decode all foods
	if err := result.All(ctx.Context(), &foods); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any food")
	}

	//Return food
	return ctx.JSON(foods)
}

func CreateFood(ctx *fiber.Ctx) error {
	//Collect store id from params
	storeId := ctx.Params("s_id")

	//Create food and menu model
	food := new(models.Food)
	menu := new(models.Menu)

	//Validate body
	utils.ParseBodyAndValidate(ctx, food)

	//Find menu by menu id and store id
	err := menuCollection.FindOne(ctx.Context(), bson.M{"menu_id": food.Menu_id, "store_id": storeId}).Decode(&menu)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Menu not found")
	}

	//Set menu name and store id
	food.Menu_Name = menu.Category
	food.Store_id = menu.Store_id

	//Set created_at and updated_at
	food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	//Create ID
	food.ID = primitive.NewObjectID()
	food.Food_id = food.ID.Hex()

	//Set price to 2 decimal places
	var num = utils.ToFixed(*food.Price, 2)
	food.Price = &num

	//Insert food
	result, err := foodCollection.InsertOne(ctx.Context(), food)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create food")
	}

	fmt.Println(result)

	//Return food
	return ctx.JSON(food)
}

func UpdateFood(ctx *fiber.Ctx) error {
	//Collect food id and store id from params
	foodId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create food model
	food := new(models.Food)

	//Find food by food id and store id
	err := foodCollection.FindOne(ctx.Context(), bson.M{"store_id": storeId, "food_id": foodId}).Decode(&food)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Food not found")
	}

	//Validate body
	utils.ParseBodyAndValidate(ctx, food)

	//Set updated_at
	food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// Create an update document with the $set operator
	update := bson.M{
		"$set": bson.M{
			"name":       food.Name,
			"price":      food.Price,
			"food_image": food.Food_Image,
			"updated_at": time.Now(),
		},
	}

	//Update food by food id and store id
	result, err := foodCollection.UpdateOne(ctx.Context(), bson.M{"food_id": foodId, "store_id": storeId}, update)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update food")
	}

	//Return result
	return ctx.JSON(result)
}

func DeleteFood(ctx *fiber.Ctx) error {
	//Collect food id and store id from params
	storeId := ctx.Params("s_id")
	foodId := ctx.Params("id")

	//Delete food by food id and store id
	result, err := foodCollection.DeleteOne(ctx.Context(), bson.M{"food_id": foodId, "store_id": storeId})

	//Return error if food not found
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete food")
	}

	//Return result
	return ctx.JSON(result)
}
