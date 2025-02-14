package handlers

import (
    "morent/database"
    "morent/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
    var cars []models.Car
    var total int64

    // Ambil parameter halaman dan batas per halaman
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    offset := (page - 1) * limit

    // Hitung total data
    database.DB.Model(&models.Car{}).Count(&total)

    // Ambil data dengan limit & offset
    result := database.DB.Limit(limit).Offset(offset).Find(&cars)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
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
