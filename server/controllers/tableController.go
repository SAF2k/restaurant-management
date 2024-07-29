package controllers

import (
	"encoding/hex"

	"time"

	"github.com/SAF2k/restaurant-management/server/helper"
	"github.com/SAF2k/restaurant-management/server/models"
	"github.com/SAF2k/restaurant-management/server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllTables(ctx *fiber.Ctx) error {

	// Collect store id from params
	storeId := ctx.Params("s_id")

	// Find all tables using store id
	tables, err := helper.GetAllTablesFromDB(storeId)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(tables)

}

func GetTable(ctx *fiber.Ctx) error {

	// Collect table id and store id from params
	table_id := ctx.Params("id")
	storeId := ctx.Params("s_id")

	// Find table using store and table ID
	table, err := helper.GetTableFromDB(table_id, storeId)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(table)

}

func CreateTable(ctx *fiber.Ctx) error {

	// Collect store ID from params
	storeId := ctx.Params("s_id")

	// Create new instance of table model
	table := new(models.Table)

	// Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, table); err != nil {
		return err
	}

	// Generate a new UUID (Universally Unique Identifier)
	tableID := uuid.New()

	// Convert UUID to string and assign it to StoreID
	table.TableID = hex.EncodeToString(tableID[:])

	// Set created_at and updated_at
	table.CreatedAt = time.Now()
	table.UpdatedAt = time.Now()

	// Set store id
	table.StoreID = storeId

	// Insert table
	if err := helper.CreateTableInDB(table); err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(table)

}

func UpdateTable(ctx *fiber.Ctx) error {

	// Collect table  and store ID from params
	tableID := ctx.Params("id")
	storeID := ctx.Params("s_id")

	// Create table struct
	table := &models.Table{}

	updateValues := map[string]interface{}{
		"number_of_guests": table.NumberOfGuests,
		"table_number":     table.TableNumber,
		"updated_at":       time.Now(),
	}

	// Update table
	if err := helper.UpdateTableInDB(tableID, storeID, updateValues); err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(table)

}

func DeleteTable(ctx *fiber.Ctx) error {

	// Collect table id and store ID from params
	tableID := ctx.Params("id")
	storeID := ctx.Params("s_id")

	// Delete table
	rowsAffected, err := helper.DeleteTableFromDB(tableID, storeID)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	response := map[string]interface{}{
		"table deleted successfully": rowsAffected,
	}
	return ctx.JSON(response)

}
