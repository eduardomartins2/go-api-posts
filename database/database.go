package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "go-api-posts/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.Post{}, &models.Comment{})
    DB = database
}
