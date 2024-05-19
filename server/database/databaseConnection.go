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

	//Connects to DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the Food schema
	if err := db.AutoMigrate(&models.Food{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}
	fmt.Println("Auto-migrate the Food schema" + "\t\t[OK]")

	// Auto-migrate the Invoice schema
	if err := db.AutoMigrate(&models.Invoice{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	fmt.Println("Auto-migrate the Invoice schema" + "\t\t[OK]")

	// Auto-migrate the Menu schema
	if err := db.AutoMigrate(&models.Menu{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	fmt.Println("Auto-migrate the Menu schema" + "\t\t[OK]")
	// Auto-migrate the Note schema
	if err := db.AutoMigrate(&models.Note{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	fmt.Println("Auto-migrate the Note schema" + "\t\t[OK]")
	// Auto-migrate the OrderItem schema
	if err := db.AutoMigrate(&models.OrderItem{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	fmt.Println("Auto-migrate the OrderItem schema" + "\t\t[OK]")
	// Auto-migrate the Order schema
	if err := db.AutoMigrate(&models.Order{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	fmt.Println("Auto-migrate the Order schema" + "\t\t[OK]")
	// Auto-migrate the Store schema
	if err := db.AutoMigrate(&models.Store{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	fmt.Println("Auto-migrate the Store schema" + "\t\t[OK]")

	// Auto-migrate the Table schema
	if err := db.AutoMigrate(&models.Table{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	fmt.Println("Auto-migrate the Table schema" + "\t\t[OK]")
	fmt.Println("Database connected successfully!")
	return db
}

var Client *gorm.DB = DBMariainstance()

func OpenTable(db *gorm.DB, tableName string) *gorm.DB {
	return db.Table(tableName)
}
