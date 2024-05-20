package database

import (
	"fmt"
	"log"

	"github.com/SAF2k/restaurant-management/config"
	"github.com/SAF2k/restaurant-management/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBMariainstance() *gorm.DB {
	dbUser := config.DB_USER
	dbPassword := config.DB_PASSWORD
	dbHost := config.DB_HOST
	dbPort := config.DB_PORT
	dbName := config.DB_NAME

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("MariaDB Running on ", dsn)

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate schemas
	autoMigrateSchemas(db)

	fmt.Println("Database connected successfully!")
	return db
}

func autoMigrateSchemas(db *gorm.DB) {
	schemas := []interface{}{
		&models.Food{},
		&models.Invoice{},
		&models.Menu{},
		&models.Note{},
		&models.OrderItem{},
		&models.Order{},
		&models.Store{},
		&models.Table{},
	}

	for _, schema := range schemas {
		if err := db.AutoMigrate(schema); err != nil {
			log.Fatalf("Failed to auto-migrate schema %T: %v", schema, err)
		}
		fmt.Printf("Auto-migrate schema %T\t\t[OK]\n", schema)
	}
}

var Client *gorm.DB = DBMariainstance()

func OpenTable(db *gorm.DB, tableName string) *gorm.DB {
	return db.Table(tableName)
}
