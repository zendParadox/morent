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
	}

	return r
}
