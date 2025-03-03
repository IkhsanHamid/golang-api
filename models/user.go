package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"-"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Carts []Carts `gorm:"foreignKey:UserID" json:"carts"`
}
