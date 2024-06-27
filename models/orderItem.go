package models

type OrderItem struct {
    ID        uint    `gorm:"primaryKey" json:"id"`
    OrderID   uint    `json:"order_id"`
    ProductID uint    `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
    Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
