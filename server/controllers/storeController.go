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

var storeCollection *gorm.DB = database.OpenTable(database.Client, "stores")

func GetAllStores(ctx *fiber.Ctx) error {

	//Create list of store model
	var stores []models.Store

	// Find all stores
	result := storeCollection.Find(&stores)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any store")
	}

	// Return stores
	return ctx.JSON(stores)

}

func GetStore(ctx *fiber.Ctx) error {

	//Collect id from params
	storeId := ctx.Params("id")

	//Create store model
	var store models.Store

	//Find store using store ID
	result := storeCollection.Where("store_id = ?", storeId).First(&store)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find store")
	}

	//Return store
	return ctx.JSON(store)

}

func CreateStore(ctx *fiber.Ctx) error {

	//Create new instance of store model
	store := new(models.Store)

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, store); err != nil {
		return err
	}

	// Generate a new UUID (Universally Unique Identifier)
	storeID := uuid.New()

	// Convert UUID to string and assign it to StoreID
	store.StoreID = hex.EncodeToString(storeID[:])

	// Set ID with the UUID
	store.ID = storeID

	// Set created_at and updated_at
	now := time.Now()
	store.CreatedAt = now
	store.UpdatedAt = now

	// Insert store
	if err := storeCollection.Create(store).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create store")
	}

	// Return store
	return ctx.JSON(store)

}

func UpdateStore(ctx *fiber.Ctx) error {

	//Collect store id from params
	storeID := ctx.Params("id")

	// Create store struct
	store := &models.Store{}

	// Parse body and validate
	if err := ctx.BodyParser(store); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Create update map
	update := map[string]interface{}{
		"name": store.Name,
	}

	// Update store
	result := storeCollection.Model(&models.Store{}).Where("store_id = ?", storeID).Updates(update)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update store")
	}

	// Return store
	return ctx.JSON(store)

}

func DeleteStore(ctx *fiber.Ctx) error {

	//Collect store id from params
	storeID := ctx.Params("id")

	// Delete store using store ID
	result := storeCollection.Where("store_id = ?", storeID).Delete(&models.Store{})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete store")
	}

	// Set response
	response := map[string]interface{}{
		"store deleted successfully": result.RowsAffected,
	}

	//Return response
	return ctx.JSON(response)

}
