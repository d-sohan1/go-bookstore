package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/d-sohan1/go-bookstore/pkg/models"
	"github.com/d-sohan1/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	fetchedBooks := models.GetAllBooks()
	res, _ := json.Marshal(fetchedBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId")
	}
	fmt.Println(Id)
	fetchedBook, _ := models.GetBookById(Id)
	res, _ := json.Marshal(fetchedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	createdBook := book.CreateBook()
	res, _ := json.Marshal(createdBook)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId")
	}
	deletedBook := models.DeleteBookById(Id)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	params := mux.Vars(r)
	Id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId")
	}
	book, db := models.GetBookById(Id)
	if newBook.Name != "" {
		book.Name = newBook.Name
	}
	if newBook.Author != "" {
		book.Author = newBook.Author
	}
	if newBook.Publication != "" {
		book.Publication = newBook.Publication
	}
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
