package controllers

import (
	"encoding/hex"
	"time"

	"github.com/SAF2k/restaurant-management/server/helper"
	"github.com/SAF2k/restaurant-management/server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllStores(ctx *fiber.Ctx) error {

	// Find all store
	stores, err := helper.GetAllStoresFromDB()
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(stores)

}

func GetStore(ctx *fiber.Ctx) error {

	// Collect store ID from params
	storeId := ctx.Params("id")

	// Find store using store ID
	store, err := helper.GetStoreFromDB(storeId)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(store)

}

func CreateStore(ctx *fiber.Ctx) error {

	store := new(models.Store)

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
	if err := helper.CreateStoreInDB(store); err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(store)

}

func UpdateStore(ctx *fiber.Ctx) error {

	// Collect store ID from params
	storeID := ctx.Params("id")

	store := &models.Store{}

	update := map[string]interface{}{
		"name": store.Name,
	}

	// Update store
	if err := helper.UpdateStoreInDB(storeID, update); err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(store)

}

func DeleteStore(ctx *fiber.Ctx) error {

	// Collect store ID from params
	storeID := ctx.Params("id")

	// Delete store
	rowsAffected, err := helper.DeleteStoreFromDB(storeID)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	response := map[string]interface{}{
		"store deleted successfully": rowsAffected,
	}
	return ctx.JSON(response)

}
