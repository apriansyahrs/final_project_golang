package database

import (
	"final_project_golang/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env:", err)
	}
}

func StartDB() {
	loadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	if isDebugMode() {
		runMigrations()
	}
}

func GetDB() *gorm.DB {
	loadEnv()

	if isDebugMode() {
		return db.Debug()
	}

	return db
}

func isDebugMode() bool {
	debugModeStr := os.Getenv("DEBUG_MODE")
	debugMode, err := strconv.ParseBool(debugModeStr)
	if err != nil {
		return false
	}
	return debugMode
}

func runMigrations() {
	err := db.AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}
}
