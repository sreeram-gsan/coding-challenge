package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Creating a variable to store DB connection.
var DB *gorm.DB = nil

// Singleton pattern for DB connection to prevent duplicate DB connections.
func getDBConnection() *gorm.DB {
	if DB == nil {
		initializeDB()
	}
	return DB
}

// Initialize DB connection and create inventory DB if needed.
func initializeDB() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connection established")

	// Create inventory table if needed.
	createInventoryTable(db)

	// Set DB global variable.
	DB = db
}

// Create inventory table.
func createInventoryTable(DB *gorm.DB) {
	query := `CREATE TABLE IF NOT EXISTS inventory(id BIGINT primary key auto_increment, name varchar(255), 
        quantity INT, unit_price FLOAT, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	res := DB.Exec(query)
	fmt.Println(res)
}
