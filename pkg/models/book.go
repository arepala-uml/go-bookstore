package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// var db *gorm.DB
var DB *gorm.DB

// gorm.model gives us a structure to help us store something in the database
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Book) CreateBook() *Book {

	DB.NewRecord(b) // initializing the new record
	DB.Create(&b)   // creating a book
	return b
}

func GetAllBooks() []Book {

	books := make([]Book, 0)
	DB.Find(&books)
	return books
}

func GetBookById(Id int) (*Book, *gorm.DB) {

	var getBook Book

	db := DB.Where("ID=?", Id).Find(&getBook) // running where command in mysql
	return &getBook, db                       // Returning the book found and db variable as well which we created of gorm type
}

func DeleteBook(Id int) Book {
	var book Book
	fmt.Println(Id)
	DB.Where("ID=?", Id).Delete(&book) // Find the book by id and then delete the book
	return book
}
