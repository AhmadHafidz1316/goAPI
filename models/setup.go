package models

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    DB.AutoMigrate(&Product{})
    DB.AutoMigrate(&User{})
}
