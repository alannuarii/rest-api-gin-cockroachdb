package models

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    sslMode := os.Getenv("DB_SSLMODE")

    dsn := "user=" + username + " password=" + password + " host=" + host + " port=" + port + " dbname=" + dbName + " sslmode=" + sslMode

    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    err = database.AutoMigrate(&Product{})
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }

    DB = database
}
