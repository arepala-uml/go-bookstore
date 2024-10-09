package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arepala-uml/go-bookstore/pkg/models"
	"github.com/arepala-uml/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {

	newBooks := models.GetAllBooks() // find all the books and get the list of books

	log.Info("Got Request for fetching all books")
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	bookId := params["bookId"]
	log.Printf("Got Request for fetching book details with id %s", bookId)
	//Id := string(bookId)
	Id, err := strconv.Atoi(bookId)
	if err != nil {
		log.Errorf("Error while parsing: %v", err)
	}

	bookDetails, _ := models.GetBookById(Id)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	log.Info("Got request for creating the book with details such as %v", CreateBook)
	b := CreateBook.CreateBook()
	//w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	bookId := params["bookId"]
	log.Info(bookId)
	Id, err := strconv.Atoi(bookId)
	if err != nil {
		log.Errorf("Error while parsing bookId: %v", err)
	}
	DeleteBook := models.DeleteBook(Id)
	log.Info("Delete Book: %v", DeleteBook)
	res, _ := json.Marshal(DeleteBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.Atoi(bookId)
	if err != nil {
		log.Errorf("Error while parsing bookId: %v", err)
	}

	bookDetails, db := models.GetBookById(Id)

	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}

	// Save it in the database
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
