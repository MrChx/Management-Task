package main

import (
	"fmt"
	"log"
	"management-task/config"
	"management-task/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	fmt.Println("Menyambungkan ke database...")
	db := config.DatabaseConnection()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mengambil objek DB: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	} else {
		fmt.Println("Koneksi ke database berhasil.")
	}

	db.AutoMigrate(&models.User{}, &models.Task{})
	config.CreateOwnerAccount(db)

	server := host + ":" + port

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Test Connection"})
	})

	fmt.Printf("Server berjalan di http://%s\n", server)
	router.Run(server)
}
