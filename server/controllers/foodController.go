package controllers

import (
	"encoding/hex"
	"time"

	"github.com/SAF2k/restaurant-management/database"
	"github.com/SAF2k/restaurant-management/models"
	"github.com/SAF2k/restaurant-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var foodCollection *gorm.DB = database.OpenTable(database.Client, "foods")

func GetAllFood(ctx *fiber.Ctx) error {

	//Collect store id from params
	storeId := ctx.Params("id")

	//Create list of food model
	var food []models.Food

	//Find all foods using store ID
	result := foodCollection.Where("store_id = ?", storeId).Find(&food)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any food")
	}

	//Return foods
	return ctx.JSON(food)
}

func GetFood(ctx *fiber.Ctx) error {

	//Collect food and store ID from params
	foodId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create food model
	var food models.Food

	//Find food using store and food ID
	result := foodCollection.Where("store_id = ? AND food_id = ?", storeId, foodId).First(&food)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find food")
	}

	//Return food
	return ctx.JSON(food)
}

func GetFoodByMenu(ctx *fiber.Ctx) error {

	//Collect menu and store ID from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create food model
	var food models.Food

	//Find food using store and menu ID
	result := foodCollection.Where("store_id = ? AND menu_id = ?", storeId, menuId).First(&food)

	// result, err := foodCollection.Find(ctx.Context(), bson.M{"menu_id": menuId, "store_id": storeId})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any food")
	}

	//Return food
	return ctx.JSON(food)
}

func CreateFood(ctx *fiber.Ctx) error {

	// Collect store id from params
	// storeId := ctx.Params("s_id")

	//Create new instance of food and menu model
	food := new(models.Food)
	menu := new(models.Menu)

	//Validate body
	if err := utils.ParseBodyAndValidate(ctx, food); err != nil {
		return err
	}

	//Set menu name and store id
	food.MenuName = menu.Category
	food.StoreID = menu.Store_id

	//Set created_at and updated_at
	food.CreatedAt = time.Now()
	food.UpdatedAt = time.Now()

	// Generate a new UUID (Universally Unique Identifier)
	foodId := uuid.New()

	// Convert UUID to string and assign it to StoreID
	food.FoodID = hex.EncodeToString(foodId[:])

	//Set price to 2 decimal places
	var num = utils.ToFixed(food.Price, 2)
	food.Price = num

	//Insert food
	if err := foodCollection.Create(food).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create food")
	}

	//Return food
	return ctx.JSON(food)
}

func UpdateFood(ctx *fiber.Ctx) error {

	//Collect food and store ID from params
	foodId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create food struct
	food := &models.Food{}

	// Find if the food is present using table and store ID
	if err := foodCollection.Where("store_id = ? AND food_id = ?", storeId, foodId).First(food).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Food not found")
	}

	//Validate body
	if err := utils.ParseBodyAndValidate(ctx, food); err != nil {
		return err
	}

	// Create update map
	update := map[string]interface{}{
		"name":       food.Name,
		"price":      food.Price,
		"food_image": food.FoodImage,
		"updated_at": time.Now(),
	}

	// Update food
	result := foodCollection.Model(&models.Food{}).Where("store_id = ? AND food_id = ?", storeId, foodId).Updates(update)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update food")
	}

	// Return food
	return ctx.JSON(food)

}

func DeleteFood(ctx *fiber.Ctx) error {

	//Collect food and store ID from params
	storeId := ctx.Params("s_id")
	foodId := ctx.Params("id")

	//Find if the food is present using store and food ID
	if err := foodCollection.Where("store_id = ? AND food_id = ? ", storeId, foodId).First(&models.Food{}).Error; err != nil {
		//Return error if food not found
		return fiber.NewError(fiber.StatusInternalServerError, "Food not found")
	}

	//Delete food using store and food
	result := foodCollection.Where("store_id = ? AND food_id = ? ", storeId, foodId).Delete(&models.Food{})
	if result.Error != nil {
		//Return error if food was not deleted
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete food")
	}

	// Set response
	response := map[string]interface{}{
		"food deleted successfully": result.RowsAffected,
	}

	//Return response
	return ctx.JSON(response)

}
