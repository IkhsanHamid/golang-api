package models

import "gorm.io/gorm"

type ProductCategory struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Product struct {
    gorm.Model
    Name           string           `json:"name"`
    Description    string           `json:"description"`
    Price          float64          `json:"price"`
    CategoryID     uint             `json:"category_id"`
    Category       ProductCategory `gorm:"foreignkey:CategoryID" json:"category"`
    StockQuantity  int              `json:"stock_quantity"`
    CartItems []CartItem `gorm:"foreignKey:ProductID" json:"CartItem"`
}
