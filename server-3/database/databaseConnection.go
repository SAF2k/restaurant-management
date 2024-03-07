package database

import (
	"fmt"
	"log"
	"restaurant-management/server-2/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInstance() *gorm.DB {
	MariaDB := config.MARIADB_URI
	fmt.Println("MariaDB Running on ", MariaDB)

	// Open a connection to the database
	db, err := gorm.Open(mysql.Open(MariaDB), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to check if the connection is alive
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Set the maximum idle connections
	sqlDB.SetMaxIdleConns(10)

	// Set the maximum open connections
	sqlDB.SetMaxOpenConns(100)

	// Set the maximum lifetime of a connection
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connected to MariaDB!")

	return db
}

var DB *gorm.DB = DBInstance()
