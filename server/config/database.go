package config

import (
	"fmt"
	"kopikasir-backend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase(){
	// load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// mengambil data dari env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// membuat DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbUser,dbPass,dbHost,dbName)
	
	// mengkoneksikan ke MySQL
	database,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		// panic("Gagal koneksi ke database!")
		log.Fatalf("Gagal koneksi ke database. DSN=%s | err=%v",dsn,err)
	}

	// auto migrate ( membuat tabel otomatis )
	if err := database.AutoMigrate(
		&models.User{},
	); err != nil{
		log.Fatalf("Gagal koneksi DB: %v", err)
	}

	// memberikan informasi berhasil terhubung ke database dan berhasil migrate
	DB = database
	fmt.Println("🚀 Database Connected & Migrated!")

}