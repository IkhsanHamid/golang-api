package config

import (
	"golang-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=ikhsan1802 dbname=golang-api port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.User{}, &models.Product{}, &models.Carts{}, &models.CartItem{}, &models.ProductCategory{})

    DB = database
}
