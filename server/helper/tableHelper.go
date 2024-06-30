package helper

import (
	"errors"

	"github.com/SAF2k/restaurant-management/server/database"
	"github.com/SAF2k/restaurant-management/server/models"

	"gorm.io/gorm"
)

// Custom errors
var (
	ErrConnectingDB_TableHelper   = errors.New("could not connect to the database")
	ErrRecordNotFound_TableHelper = errors.New("record not found")
)

func GetAllTablesFromDB(storeId string) ([]models.Table, error) {

	//Create list of table model
	var table []models.Table

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Fetch all table with store ID
	result := db.Where("store_id = ?", storeId).Find(&table)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_TableHelper
		}
		return nil, ErrConnectingDB_TableHelper
	}
	return table, nil

}

func GetTableFromDB(tableId, storeId string) (*models.Table, error) {

	var table models.Table

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Fetch all table items for the given store ID
	result := db.Where("store_id = ? AND table_id = ?", storeId, tableId).First(&table)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound_TableHelper
		}
		return nil, ErrConnectingDB_TableHelper
	}
	return &table, nil

}

func CreateTableInDB(table *models.Table) error {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if table already exists before inserting
	var existingtable models.Table
	db.Where("table_id =?", table.TableID).First(&existingtable)

	if existingtable.TableID != "" {
		return errors.New("table ID already exists")
	}

	// If not, insert the new table
	result := db.Create(table)
	if result.Error != nil {
		return ErrConnectingDB_TableHelper
	}

	return nil

}

func UpdateTableInDB(tableId, storeId string, update map[string]interface{}) error {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if table already exists before inserting
	var existingtable models.Table
	db.Where("store_id = ? AND table_id = ?", storeId, tableId).First(existingtable)

	if existingtable.TableID != "" {
		// Update if the table is present
		result := db.Model(&models.Table{}).Where("store_id = ? AND table_id = ?", storeId, tableId).Updates(update)
		if result.Error != nil {
			return ErrConnectingDB_TableHelper
		}
		return nil
	}

	return errors.New("table ID doesn't exists")
}

func DeleteTableFromDB(tableId, storeId string) (int64, error) {

	db := database.OpenDB()
	defer database.CloseDB(db)

	// Check if table exists before deleting
	var existingtable models.Table
	db.Where("store_id = ? AND table_id = ?", storeId, tableId).First(existingtable)
	if existingtable.TableID != "" {
		// Delete the store with the given store ID
		result := db.Where("store_id = ? AND table_id = ?", storeId, tableId).Delete(&models.Table{})
		if result.Error != nil {
			return 0, ErrConnectingDB_StoreHelper
		}
		return result.RowsAffected, nil
	}

	return 0, errors.New("table ID doesn't exists")

}
