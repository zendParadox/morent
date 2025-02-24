package models

import "time"

type Review struct {
    ID         int       `json:"id" gorm:"primaryKey"`
    ProductID  int       `json:"product_id"` // Foreign key ke tabel Car
    UserID     uint64    `json:"user_id" gorm:"type:bigint UNSIGNED"`
    Rating     float64   `json:"rating" gorm:"type:decimal(10,1)"`
    Comment    string    `json:"comment"`
    CreatedAt  time.Time `json:"created_at"`
}