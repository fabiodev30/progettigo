package models

import (
	"bookstore/dbcontroller"

	"gorm.io/gorm"
)

type Book struct {
	ID            int `gorm:"primaryKey"`
	Nome          string
	Autore        string
	Pubblicazione string
}

var db *gorm.DB

func init() {
	dbcontroller.Connect()
	db = dbcontroller.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
