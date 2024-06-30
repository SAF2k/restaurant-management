package controllers

import (
	"encoding/hex"

	"time"

	"github.com/SAF2k/restaurant-management/server/helper"
	"github.com/SAF2k/restaurant-management/server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllMenus(ctx *fiber.Ctx) error {

	// Collect store ID from params
	storeId := ctx.Params("id")

	// Find all menu using store ID
	menus, err := helper.GetAllMenusFromDB(storeId)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(menus)

}

func GetMenu(ctx *fiber.Ctx) error {

	// Collect menu and store ID from params
	storeId := ctx.Params("s_id")
	menuId := ctx.Params("id")

	// Find menu using store and menu ID
	menus, err := helper.GetMenuFromDB(menuId, storeId)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(menus)

}

func CreateMenu(ctx *fiber.Ctx) error {

	// Collect store id from params
	storeId := ctx.Params("s_id")

	// Create new instance of the menu model
	menu := new(models.Menu)

	// Generate a new UUID (Universally Unique Identifier)
	menuID := uuid.New()

	// Convert UUID to string and assign it to StoreID
	menu.Menu_id = hex.EncodeToString(menuID[:])

	// Set ID with the UUID
	menu.ID = menuID

	// Create a new menu model
	menuModel := models.Menu{
		ID:         menu.ID,
		Name:       menu.Name,
		Category:   menu.Category,
		Menu_id:    menu.Menu_id,
		Store_id:   &storeId,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	// Insert menu
	if err := helper.CreateMenuInDB(&menuModel); err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(menu)

}

func UpdateMenu(ctx *fiber.Ctx) error {

	// Collect menu and store ID from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	// Create menu struct
	menu := &models.Menu{}

	update := map[string]interface{}{
		"name":       menu.Name,
		"category":   menu.Category,
		"updated_at": time.Now(),
	}

	// Update menu
	if err := helper.UpdateMenuInDB(menuId, storeId, update); err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(menu)

}

func DeleteMenu(ctx *fiber.Ctx) error {

	// Collect menu and store ID from params
	menuId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	// Delete menu using menu and store ID
	rowsAffected, err := helper.DeleteMenuFromDB(menuId, storeId)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	response := map[string]interface{}{
		"menu deleted successfully": rowsAffected,
	}
	return ctx.JSON(response)

}
