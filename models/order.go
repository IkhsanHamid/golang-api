package models

import "time"

type Order struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    UserID     uint      `json:"user_id"`
    TotalPrice float64   `json:"total_price"`
    Status     string    `json:"status"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}
