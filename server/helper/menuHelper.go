package helper

import (
	"errors"

	"github.com/SAF2k/restaurant-management/server/database"
	"github.com/SAF2k/restaurant-management/server/models"

	"gorm.io/gorm"
)

// Custom errors
var (
	ErrConnectingDB_MenuHelper   = errors.New("could not connect to the database")
	ErrRecordNotFound_MenuHelper = errors.New("record not found")
)

func GetAllMenusFromDB(storeId string) ([]models.Menu, error) {

	// Create list of menu model
	var menus []models.Menu

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Fetch all menu
	result := db.Find(&menus)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_MenuHelper
		}
		return nil, ErrConnectingDB_MenuHelper
	}
	return menus, nil

}

func GetMenuFromDB(menuId, storeId string) (*models.Menu, error) {

	var menu models.Menu

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Fetch all menu items for the given store ID
	result := db.Where("store_id = ? AND menu_id = ?", storeId, menuId).First(&menu)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_MenuHelper
		}
		return nil, ErrConnectingDB_MenuHelper
	}
	return &menu, nil

}

func CreateMenuInDB(menu *models.Menu) error {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if menu already exists before inserting
	var existingMenu models.Menu
	db.Where("menu_id =?", menu.Menu_id).First(&existingMenu)

	if existingMenu.Menu_id != "" {
		return errors.New("menu ID already exists")
	}

	// If not, insert the new menu
	result := db.Create(menu)
	if result.Error != nil {
		return ErrConnectingDB_MenuHelper
	}

	return nil

}

func UpdateMenuInDB(menuId, storeId string, update map[string]interface{}) error {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if menu already exists before inserting
	var existingMenu models.Menu
	db.Where("store_id = ? AND menu_id = ?", storeId, menuId).First(&existingMenu)

	if existingMenu.Menu_id != "" {
		// Update if the menu is present
		result := db.Model(&models.Menu{}).Where("store_id = ? AND menu_id = ?", storeId, menuId).Updates(update)
		if result.Error != nil {
			return ErrConnectingDB_MenuHelper
		}
		return nil
	}

	return errors.New("menu ID doesn't exists")

}

func DeleteMenuFromDB(menuId, storeId string) (int64, error) {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if menu exists before deleting
	var existingMenu models.Menu
	db.Where("store_id = ? AND menu_id = ?", storeId, menuId)
	if existingMenu.Menu_id != "" {
		// Delete the store with the given store ID
		result := db.Where("store_id = ? AND menu_id = >", storeId, menuId).Delete(existingMenu)
		if result.Error != nil {
			return 0, ErrConnectingDB_StoreHelper
		}
		return result.RowsAffected, nil
	}

	return 0, errors.New("menu ID doesn't exists")

}
