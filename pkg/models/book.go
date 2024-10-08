package models

import (
	"fmt"

	"github.com/arepala-uml/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// gorm.model gives us a structure to help us store something in the database
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {

	db.NewRecord(b) // initializing the new record
	db.Create(&b)   // creating a book
	return b
}

func GetAllBooks() []Book {

	books := make([]Book, 0)
	db.Find(&books)
	return books
}

func GetBookById(Id int) (*Book, *gorm.DB) {

	var getBook Book

	db := db.Where("ID=?", Id).Find(&getBook) // running where command in mysql
	return &getBook, db                       // Returning the book found and db variable as well which we created of gorm type
}

func DeleteBook(Id int) Book {
	var book Book
	fmt.Println(Id)
	db.Where("ID=?", Id).Delete(&book) // Find the book by id and then delete the book
	return book
}
