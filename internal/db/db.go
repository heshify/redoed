package db

import (
	"fmt"
	"log"
	"os"

	// models "github.com/heshify/redoed/internal/models"
	"github.com/heshify/redoed/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionStr := fmt.Sprintf("host=%s user=%s	password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR: failed to initialize database %s", err)
		return
	} else {
		log.Println("INFO: Database initialized")
	}

	DB.AutoMigrate(&models.Document{})
	DB.AutoMigrate(&models.User{})
}
