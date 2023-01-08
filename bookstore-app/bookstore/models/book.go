package models

import (
	"github.com/dkr290/go-projects/bookstore-app/bookstore/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	Name        string `gorm:"name"`
	Author      string `gorm:"author"`
	Publication string `gorm:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}
