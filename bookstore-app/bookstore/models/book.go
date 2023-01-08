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

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books

}

func GetBookById(ID int64) (*Book, *gorm.DB) {

	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
