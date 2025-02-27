package handlers

import (
	"morent/database"
	"morent/models"
	"morent/utils"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// Save user to DB
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Invalid email or password"})
		return
	}

	// Cek password
	if err := user.CheckPassword(loginData.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Invalid email or password"})
		return
	}

	// Generate JWT Token
	expirationTime := time.Now().Add(time.Hour) // Token berlaku selama 1 jam
	claims := &Claims{
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "Failed to generate token"})
		return
	}

	// Response sesuai format yang diminta
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Login berhasil",
		"data": gin.H{
			"accessToken": tokenString,
			"user": gin.H{
				"userId":    user.ID,
				"username":  user.Username,
				"email":     user.Email,
				"role":      user.Role,
				"expiresIn": 3600, // Token berlaku selama 3600 detik (1 jam)
			},
		},
	})
}

func Logout(c *gin.Context) {
	// Klien cukup menghapus token di sisi mereka
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Logout berhasil",
	})
}
