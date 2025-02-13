package routes

import (
	"github.com/gin-gonic/gin"
	"morent/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Group untuk autentikasi
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register) // Pastikan ini ada
		auth.POST("/login", handlers.Login)       // Pastikan ini ada
		r.POST("/logout", handlers.Logout)
	}
	cars := r.Group("/cars")
	{
		cars.GET("/", handlers.GetCars)
	}

	return r
}
