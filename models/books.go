package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
    ID uint `json:"id" gorm:"primary_key"`
    Title string `json:"title"`
    Author string `json:"author"`
    Rating float32 `json:"rating"`
}

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open("sqlite3", "test.db")

    if err != nil {
        panic("Failed to connect to db...")
    }
    
    database.AutoMigrate(&Book{})
    DB = database
}
