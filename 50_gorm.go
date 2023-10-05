package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	ID    uint
	Code  string
	Price uint
}

func main() {
	// Open a connection to the SQLite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Create the "products" table if it doesn't exist
	db.AutoMigrate(&Product{})

	// Create a new product
	product := Product{Code: "Laptop", Price: 1000}
	db.Create(&product)

	// Read a product by ID
	var readProduct Product
	db.First(&readProduct, 1)
	fmt.Printf("Product: %+v\n", readProduct)

	// Update a product's price
	db.Model(&readProduct).Update("Price", 1200)

	// Delete a product
	db.Delete(&readProduct)

	// Close the database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
