package models

import "time"

type Car struct {
    ID           int       `json:"id" gorm:"primaryKey"`
    Brand        string    `json:"brand"`
    Model        string    `json:"model"`
    Year         int       `json:"year"`
    LicensePlate string    `json:"license_plate"`
    PricePerDay  float64   `json:"price_per_day" gorm:"type:decimal(10,2)"`
    Status       string    `json:"status"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    ImageURL     string    `json:"image_url"`
    DeletedAt    *time.Time `json:"deleted_at"`
    AverageRating float64   `json:"average_rating" gorm:"type:decimal(3,2)"`
    Reviews      []Review  `json:"reviews" gorm:"foreignKey:ProductID;references:ID"` // Relasi ke tabel Review
}