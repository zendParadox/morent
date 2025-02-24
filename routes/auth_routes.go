package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"morent/handlers"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Tambahkan middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Sesuaikan dengan URL frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Group untuk autentikasi
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		r.POST("/logout", handlers.Logout)
	}

	// Group untuk cars
	cars := r.Group("/cars")
	{
		cars.GET("/", handlers.GetCars)
		cars.GET("/:id", handlers.GetCarByID)
	}

	return r
}
