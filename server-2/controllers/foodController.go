package controllers

import (
	"math"
	"restaurant-management/server-2/database"
	"restaurant-management/server-2/models"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

var validate = validator.New()

func GetAllFood(ctx *fiber.Ctx) error {
	result, err := foodCollection.Find(ctx.Context(), bson.M{})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	var foods []bson.M

	if err := result.All(ctx.Context(), &foods); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any food")
	}

	return ctx.JSON(foods)
}

func GetFood(ctx *fiber.Ctx) error {
	foodId := ctx.Params("id")

	var food models.Food

	err := foodCollection.FindOne(ctx.Context(), bson.M{"food_id": foodId}).Decode(&food)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Food not found")
	}

	return ctx.JSON(food)
}

func GetFoodByMenu(ctx *fiber.Ctx) error {
	menuId := ctx.Params("id")

	result, err := foodCollection.Find(ctx.Context(), bson.M{"menu_id": menuId})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	var foods []models.Food

	if err := result.All(ctx.Context(), &foods); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any food")
	}

	return ctx.JSON(foods)
}

func CreateFood(ctx *fiber.Ctx) error {
	food := new(models.Food)
	menu := new(models.Menu)

	if err := ctx.BodyParser(food); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	validationError := validate.Struct(food)
	if validationError != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	err := menuCollection.FindOne(ctx.Context(), bson.M{"menu_id": food.Menu_id}).Decode(&menu)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Menu not found")
	}

	food.Menu_Name = menu.Category
	food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.ID = primitive.NewObjectID()
	food.Food_id = food.ID.Hex()
	var num = toFixed(*food.Price, 2)
	food.Price = &num

	result, err := foodCollection.InsertOne(ctx.Context(), food)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create food")
	}

	return ctx.JSON(result)
}

func UpdateFood(ctx *fiber.Ctx) error {
	foodId := ctx.Params("id")

	food := new(models.Food)

	if err := ctx.BodyParser(food); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	validationError := validate.Struct(food)
	if validationError != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

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

	result, err := foodCollection.UpdateOne(ctx.Context(), bson.M{"food_id": foodId}, update)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update food")
	}

	return ctx.JSON(result)
}

func DeleteFood(ctx *fiber.Ctx) error {
	foodId := ctx.Params("id")

	result, err := foodCollection.DeleteOne(ctx.Context(), bson.M{"food_id": foodId})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete food")
	}

	return ctx.JSON(result)
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
