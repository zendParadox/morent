package main

import (
	"fmt"
	"log"
	"os"

	"morent/database"
	"morent/routes"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi database
	database.ConnectDB()

	// Setup router
	r := routes.SetupRouter()

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port", port)
	r.Run(":" + port)
}
