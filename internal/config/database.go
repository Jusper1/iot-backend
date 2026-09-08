package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: file .env tidak ditemukan")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}

	log.Println("Database berhasil terhubung")

	return db
}