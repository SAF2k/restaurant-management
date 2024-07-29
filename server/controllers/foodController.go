package controllers

import (
	"encoding/hex"

	"github.com/SAF2k/restaurant-management/server/utils"

	"github.com/SAF2k/restaurant-management/server/helper"
	"github.com/SAF2k/restaurant-management/server/models"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllFood(ctx *fiber.Ctx) error {

	// Collect store ID from params
	storeId := ctx.Params("id")

	// Find all food
	food, err := helper.FindAllFoodsByStoreIdFromDB(storeId)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(food)

}

func GetFood(ctx *fiber.Ctx) error {

	// Collect food and store ID from params
	foodId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	// Find food with store and food ID
	food, err := helper.FindFoodByIdAndStoreIdFromDB(foodId, storeId)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(food)

}

func GetFoodByMenu(ctx *fiber.Ctx) error {

	// Collect menu and store ID from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	// Find food with store and menu ID
	food, err := helper.FindFoodByMenuIdAndStoreIdFromDB(menuId, storeId)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(food)

}

func CreateFood(ctx *fiber.Ctx) error {

	food := new(models.Food)
	menu := new(models.Menu)

	food.MenuName = menu.Category
	food.StoreID = menu.Store_id
	food.CreatedAt = time.Now()
	food.UpdatedAt = time.Now()

	foodId := uuid.New()
	food.FoodID = hex.EncodeToString(foodId[:])
	food.Price = utils.ToFixed(food.Price, 2)

	// Insert food
	if err := helper.CreateFoodInDB(food); err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(food)

}

func UpdateFood(ctx *fiber.Ctx) error {

	// Collect food and store ID from params
	foodId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	food := &models.Food{}

	update := map[string]interface{}{
		"name":       food.Name,
		"price":      food.Price,
		"food_image": food.FoodImage,
		"updated_at": time.Now(),
	}

	// Update food
	if err := helper.UpdateFoodInDB(foodId, storeId, update); err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(food)

}

func DeleteFood(ctx *fiber.Ctx) error {

	// Collect food id and store ID from params
	storeId := ctx.Params("s_id")
	foodId := ctx.Params("id")

	// Delete food
	rowsAffected, err := helper.DeleteFoodFromDB(foodId, storeId)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	response := map[string]interface{}{
		"food deleted successfully": rowsAffected,
	}
	return ctx.JSON(response)

}
