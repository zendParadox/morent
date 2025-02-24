package handlers

import (
    "morent/database"
    "morent/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// GetCars mengambil daftar mobil dengan pagination dan rata-rata rating
func GetCars(c *gin.Context) {
    var cars []models.Car
    var total int64

    // Ambil parameter halaman dan batas per halaman
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    offset := (page - 1) * limit

    // Hitung total data
    database.DB.Model(&models.Car{}).Count(&total)

    // Ambil data mobil dengan limit & offset
    result := database.DB.Limit(limit).Offset(offset).Find(&cars)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    // Hitung rata-rata rating untuk setiap mobil
    for i, car := range cars {
        var avgRating float64
        database.DB.Model(&models.Review{}).
            Select("COALESCE(AVG(rating), 0)").
            Where("product_id = ?", car.ID).
            Scan(&avgRating)
        cars[i].AverageRating = avgRating
    }

    // Response dengan format JSON
    c.JSON(http.StatusOK, gin.H{
        "status": 200,
        "message": "Data mobil berhasil diambil",
        "data": cars,
        "pagination": gin.H{
            "current_page": page,
            "per_page": limit,
            "total_data": total,
            "total_pages": (total + int64(limit) - 1) / int64(limit),
        },
    })
}

// GetCarByID mengambil data mobil berdasarkan ID beserta review-nya
func GetCarByID(c *gin.Context) {
    carID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
        return
    }

    var car models.Car

    // Ambil data mobil berdasarkan ID dan preload reviews
    result := database.DB.Preload("Reviews", "product_id = ?", carID).First(&car, carID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
        return
    }

    // Response dengan format JSON
    c.JSON(http.StatusOK, gin.H{
        "status":  200,
        "message": "Data mobil berhasil diambil",
        "data":    car,
    })
}