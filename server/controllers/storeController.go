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

// GetAllStores - Fetch all stores
func GetAllStores(ctx *fiber.Ctx) error {
	var stores []models.Store

	result := storeCollection.Find(&stores)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	return ctx.JSON(stores)
}

// GetStore - Fetch a single store by ID
func GetStore(ctx *fiber.Ctx) error {
	storeID := ctx.Params("id")
	var store models.Store

	result := storeCollection.Where("store_id = ?", storeID).First(&store)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "Store not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find store")
	}

	return ctx.JSON(store)
}

// CreateStore - Create a new store
func CreateStore(ctx *fiber.Ctx) error {
	store := new(models.Store)

	if err := utils.ParseBodyAndValidate(ctx, store); err != nil {
		return err
	}

	storeID := uuid.New()
	store.StoreID = hex.EncodeToString(storeID[:])
	store.ID = storeID
	now := time.Now()
	store.CreatedAt = now
	store.UpdatedAt = now

	if err := storeCollection.Create(store).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create store")
	}

	return ctx.JSON(store)
}

// UpdateStore - Update an existing store
func UpdateStore(ctx *fiber.Ctx) error {
	storeID := ctx.Params("id")
	var store models.Store

	if err := ctx.BodyParser(&store); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	update := map[string]interface{}{
		"name":       store.Name,
		"updated_at": time.Now(),
	}

	result := storeCollection.Model(&models.Store{}).Where("store_id = ?", storeID).Updates(update)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update store")
	}

	return ctx.JSON(store)
}

// DeleteStore - Delete a store by ID
func DeleteStore(ctx *fiber.Ctx) error {
	storeID := ctx.Params("id")

	result := storeCollection.Where("store_id = ?", storeID).Delete(&models.Store{})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete store")
	}

	response := map[string]interface{}{
		"store deleted successfully": result.RowsAffected,
	}

	return ctx.JSON(response)
}
