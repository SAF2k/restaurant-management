package database

import (
	"fmt"
	"log"

	"github.com/SAF2k/restaurant-management/server/config"
	"github.com/SAF2k/restaurant-management/server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbUser     = config.DB_USER
	dbPassword = config.DB_PASSWORD
	dbHost     = config.DB_HOST
	dbPort     = config.DB_PORT
	dbName     = config.DB_NAME

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	Client *gorm.DB
)

func init() {
	Client = DBMariainstance()
}

// DBMariainstance initializes and returns a database connection
func DBMariainstance() *gorm.DB {
	fmt.Println("MariaDB Running on ", dsn)

	// Connects to DB
	db := OpenDB()

	// Auto-migrate schemas
	modelsToMigrate := []interface{}{
		&models.Food{},
		&models.Invoice{},
		&models.Menu{},
		&models.Note{},
		&models.OrderItem{},
		&models.Order{},
		&models.Store{},
		&models.Table{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to auto-migrate database schema for %T: %v", model, err)
		}
		fmt.Printf("Auto-migrate the %T schema\t\t[OK]\n", model)
	}

	fmt.Println("Connected to DB and Auto-migrated all the tables successfully!")
	return db
}

func OpenDB() *gorm.DB {
	// Connects to DB
	cred, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return cred
}

// CloseDB closes the database connection
func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// OpenTable returns a DB instance for the specified table
func OpenTable(tableName string) *gorm.DB {
	return Client.Table(tableName)
}
