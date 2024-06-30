package helper

import (
	"errors"

	"github.com/SAF2k/restaurant-management/server/database"
	"github.com/SAF2k/restaurant-management/server/models"

	"gorm.io/gorm"
)

// Custom errors
var (
	ErrConnectingDB_FoodHelper   = errors.New("could not connect to the database")
	ErrRecordNotFound_FoodHelper = errors.New("record not found")
)

func FindAllFoodsByStoreIdFromDB(storeId string) ([]models.Food, error) {

	var food []models.Food

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Fetch all food items with store ID
	result := db.Where("store_id = ?", storeId).Find(&food)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_FoodHelper
		}
		return nil, ErrConnectingDB_FoodHelper
	}
	return food, nil

}

func FindFoodByIdAndStoreIdFromDB(foodId, storeId string) (*models.Food, error) {

	var food models.Food

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Fetch all food items with store and food ID
	result := db.Where("store_id = ? AND food_id = ?", storeId, foodId).First(&food)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_FoodHelper
		}
		return nil, ErrConnectingDB_FoodHelper
	}
	return &food, nil

}

func FindFoodByMenuIdAndStoreIdFromDB(menuId, storeId string) (*models.Food, error) {

	var food models.Food

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Fetch all food items with store and menu ID
	result := db.Where("store_id = ? AND menu_id = ?", storeId, menuId).First(&food)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_FoodHelper
		}
		return nil, ErrConnectingDB_FoodHelper
	}
	return &food, nil

}

func CreateFoodInDB(food *models.Food) error {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if food already exists before inserting
	var existingFood models.Food
	db.Where("food_id =?", food.FoodID).First(&existingFood)

	if existingFood.FoodID != "" {
		return errors.New("food ID already exists")
	}

	// If not, insert the new food
	result := db.Create(food)
	if result.Error != nil {
		return ErrConnectingDB_FoodHelper
	}

	return nil

}

func UpdateFoodInDB(foodId, storeId string, update map[string]interface{}) error {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if food exists before updating
	var existingFood *models.Food
	existingFood, err := FindFoodByIdAndStoreIdFromDB(foodId, storeId)
	if err != nil {
		return ErrRecordNotFound_FoodHelper
	}

	if existingFood.FoodID != "" {
		// Update if the food is present
		result := db.Model(&models.Food{}).Where("store_id = ? AND food_id = ?", storeId, foodId).Updates(update)
		if result.Error != nil {
			return ErrConnectingDB_FoodHelper
		}
		return nil
	}

	return errors.New("food ID doesn't exists")

}

func DeleteFoodFromDB(foodId, storeId string) (int64, error) {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if food exists before deleting
	var existingFood models.Food
	db.Where("store_id = ? AND food_id = ? ", storeId, foodId).First(existingFood)
	if existingFood.FoodID != "" {
		// Delete the food with the given store and food ID
		result := db.Where("store_id = ? AND food_id = ? ", storeId, foodId).Delete(existingFood)
		if result.Error != nil {
			return 0, ErrConnectingDB_StoreHelper
		}
		return result.RowsAffected, nil
	}

	return 0, errors.New("food ID doesn't exists")

}
