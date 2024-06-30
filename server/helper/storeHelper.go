package helper

import (
	"errors"

	"github.com/SAF2k/restaurant-management/server/database"
	"github.com/SAF2k/restaurant-management/server/models"

	"gorm.io/gorm"
)

// Custom errors
var (
	ErrConnectingDB_StoreHelper   = errors.New("could not connect to the database")
	ErrRecordNotFound_StoreHelper = errors.New("record not found")
)

func GetAllStoresFromDB() ([]models.Store, error) {

	var stores []models.Store

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Query to get all stores
	result := db.Find(&stores)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_StoreHelper
		}
		return nil, ErrConnectingDB_StoreHelper
	}
	return stores, nil

}

func GetStoreFromDB(storeId string) (*models.Store, error) {

	var store models.Store

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Query to get a specific store by store ID
	result := db.Where("store_id = ?", storeId).First(&store)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_StoreHelper
		}
		return nil, ErrConnectingDB_StoreHelper
	}
	return &store, nil

}

func CreateStoreInDB(store *models.Store) error {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Before inserting, check if the store ID already exists
	var existingStore models.Store
	db.Where("store_id =?", store.StoreID).First(&existingStore)
	if existingStore.StoreID != "" {
		return errors.New("store ID already exists")
	}

	// If not, insert the new store
	result := db.Create(store)
	if result.Error != nil {
		return ErrConnectingDB_StoreHelper
	}

	return nil

}

func UpdateStoreInDB(storeID string, update map[string]interface{}) error {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Update the store with the given store ID
	result := db.Model(&models.Store{}).Where("store_id = ?", storeID).Updates(update)
	if result.Error != nil {
		return ErrConnectingDB_StoreHelper
	}
	return nil

}

func DeleteStoreFromDB(storeID string) (int64, error) {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Delete the store with the given store ID
	result := db.Where("store_id = ?", storeID).Delete(&models.Store{})
	if result.Error != nil {
		return 0, ErrConnectingDB_StoreHelper
	}
	return result.RowsAffected, nil

}
