package models

import "gorm.io/gorm"

type Car struct {
    gorm.Model
    Name       string  `json:"name"`
    Brand      string  `json:"brand"`
    Year       int     `json:"year"`
    Price      float64 `json:"price"`
    ImageURL   string  `json:"image_url"`
}
