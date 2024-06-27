package models

import "gorm.io/gorm"

type Carts struct {
    gorm.Model
    ID        uint `gorm:"primaryKey" json:"id"`
    UserID    uint `json:"user_id"`
    User      User `gorm:"foreignKey:UserID" json:"user"`
    Items     []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

type CartItem struct {
    gorm.Model
    ID        uint `gorm:"primaryKey" json:"id"`
    CartID    uint `json:"cart_id"`
    Cart      Carts `gorm:"foreignKey:CartID" json:"-"`
    ProductID uint `json:"product_id"`
    Product   Product `gorm:"foreignKey:ProductID" json:"product"`
    Quantity  int  `json:"quantity"`
}
