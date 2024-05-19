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

var tableCollection *gorm.DB = database.OpenTable(database.Client, "tables")

func GetAllTables(ctx *fiber.Ctx) error {

	//Collect store id from params
	storeId := ctx.Params("s_id")

	//Create list of table model
	var tables []models.Table

	//Find all tables using store id
	result := tableCollection.Where("store_id = ?", storeId).Find(&tables)

	// result := tableCollection.Find(&tables)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any table")
	}

	//Return table
	return ctx.JSON(tables)

}

func GetTable(ctx *fiber.Ctx) error {

	//Collect table id and store id from params
	table_id := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create Store model
	var table models.Table

	//Find table using store and table ID
	result := tableCollection.Where("store_id = ? AND table_id = ?", storeId, table_id).First(&table)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find table")
	}

	//Return table
	return ctx.JSON(table)

}

func CreateTable(ctx *fiber.Ctx) error {

	//Collect store id from params
	storeId := ctx.Params("s_id")

	//Create new instance of table model
	table := new(models.Table)

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, table); err != nil {
		return err
	}

	// Generate a new UUID (Universally Unique Identifier)
	tableID := uuid.New()

	// Convert UUID to string and assign it to StoreID
	table.TableID = hex.EncodeToString(tableID[:])

	//Set created_at and updated_at
	table.CreatedAt = time.Now()
	table.UpdatedAt = time.Now()

	//Set store id
	table.StoreID = storeId

	//Insert table
	if err := tableCollection.Create(table).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create table")
	}

	//Return table
	return ctx.JSON(table)

}

func UpdateTable(ctx *fiber.Ctx) error {

	//Collect table id and store id from params
	tableID := ctx.Params("id")
	storeID := ctx.Params("s_id")

	// Create table struct
	table := &models.Table{}

	// Find if table present using table and store ID
	if err := tableCollection.Where("store_id = ? AND table_id = ?", storeID, tableID).First(table).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Table not found")
	}

	// Parse body and validate (assuming you have utils.ParseBodyAndValidate function)
	if err := utils.ParseBodyAndValidate(ctx, table); err != nil {
		return err
	}

	// Update the table attributes
	updateValues := map[string]interface{}{
		"number_of_guests": table.NumberOfGuests,
		"table_number":     table.TableNumber,
		"updated_at":       time.Now(),
	}

	//Update table
	if err := tableCollection.Model(table).Where("store_id = ? AND table_id = ?", storeID, tableID).Updates(updateValues).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update table")
	}

	// Return table
	return ctx.JSON(table)
}

func DeleteTable(ctx *fiber.Ctx) error {
	//Collect table id and store id from params
	tableID := ctx.Params("id")
	storeID := ctx.Params("s_id")

	// Create table struct
	table := &models.Table{}

	// Find if table is present using table and store ID
	if err := tableCollection.Where("store_id = ? AND table_id = ?", storeID, tableID).First(table).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Table not found")
	}

	// Perform the delete using GORM
	result := tableCollection.Where("store_id = ? AND table_id = ?", storeID, tableID).Delete(table)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete table")
	}
	// Set response
	response := map[string]interface{}{
		"Table deleted successfully": result.RowsAffected,
	}
	// Return response
	return ctx.JSON(response)
}
