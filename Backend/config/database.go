package config

import (
	"fmt"
	"log"
	"os"

	"management-task/models"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return database
}

func CreateOwnerAccount(db *gorm.DB) {
	email := os.Getenv("ADMIN_EMAIL")
	password := os.Getenv("ADMIN_PASSWORD")

	hashedPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	owner := models.User{
		Role:     "Admin",
		Name:     "Owner",
		Password: string(hashedPasswordBytes),
		Email:    email,
	}

	if db.Where("email = ?", owner.Email).First(&owner).RowsAffected == 0 {
		db.Create(&owner)
	} else {
		fmt.Println("owner exist")
	}
}
