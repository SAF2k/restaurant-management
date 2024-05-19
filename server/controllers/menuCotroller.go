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

var menuCollection *gorm.DB = database.OpenTable(database.Client, "menus")

func GetAllMenus(ctx *fiber.Ctx) error {

	//Create list of menu model
	var menus []models.Menu

	// Find all menu
	result := menuCollection.Find(&menus)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any menu")
	}

	//Return menu
	return ctx.JSON(menus)

}

func GetMenu(ctx *fiber.Ctx) error {

	//Collect menu and store ID from params
	storeId := ctx.Params("s_id")
	menuId := ctx.Params("id")

	//Create menu model
	var menus models.Menu

	//Find menu using store and menu ID
	result := storeCollection.Where("store_id = ? AND menu_id = ?", storeId, menuId).First(&menus)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find menu")
	}

	//Return menu
	return ctx.JSON(menus)

}

func CreateMenu(ctx *fiber.Ctx) error {

	//Collect store id from params
	storeId := ctx.Params("s_id")

	//Create new instance of the menu model
	menus := new(models.Menu)

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, menus); err != nil {
		return err
	}

	// Generate a new UUID (Universally Unique Identifier)
	menuID := uuid.New()

	// Convert UUID to string and assign it to StoreID
	menus.Menu_id = hex.EncodeToString(menuID[:])

	// Set ID with the UUID
	menus.ID = menuID

	//Create a new menu model
	menuModel := models.Menu{
		ID:         menus.ID,
		Name:       menus.Name,
		Category:   menus.Category,
		Menu_id:    menus.Menu_id,
		Store_id:   &storeId,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	// Insert menu
	if err := menuCollection.Create(menuModel).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create menu")
	}

	// Return menu
	return ctx.JSON(menus)

}

func UpdateMenu(ctx *fiber.Ctx) error {

	//Collect menu and store ID from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create menu struct
	menu := &models.Menu{}

	//Find menu by menu and store ID
	err := menuCollection.Where("store_id = ? AND menu_id = ?", storeId, menuId).First(&menu)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Menu not found")
	}

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, menu); err != nil {
		return err
	}

	update := map[string]interface{}{
		"name":       menu.Name,
		"category":   menu.Category,
		"updated_at": time.Now(),
	}

	// Update menu
	result := menuCollection.Model(&models.Menu{}).Where("store_id = ? AND table_id = ?", storeId, menuId).Updates(update)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update menu")
	}

	// Return menu
	return ctx.JSON(menu)

}

func DeleteMenu(ctx *fiber.Ctx) error {

	//Collect menu and store ID from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	// Create menu struct
	menu := &models.Menu{}

	// Find if menu is present using menu and store ID
	if err := menuCollection.Where("store_id = ? AND menu_id = ?", storeId, menuId).First(menu).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Menu not found")
	}

	//Delete menu using menu and store ID
	result := menuCollection.Where("store_id = ? AND menu_id = >", storeId, menuId).Delete(menu)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete menu")
	}

	// Set response
	response := map[string]interface{}{
		"Menu deleted successfully": result.RowsAffected,
	}

	// Return response
	return ctx.JSON(response)

}
