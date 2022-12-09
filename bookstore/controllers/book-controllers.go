package controllers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	creareBook := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&creareBook)
	if err != nil {
		return
	}
	b := models.CreateBook(&creareBook)
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookbyId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookid"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		return
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(&bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	err := json.NewDecoder(r.Body).Decode(&updateBook)
	if err != nil {
		return
	}
	params := mux.Vars(r)
	bookID := params["bookid"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		return
	}
	bookDetails, db := models.GetBookById(id)
	if updateBook.Nome != "" {
		bookDetails.Nome = updateBook.Nome
	}
	if updateBook.Autore != "" {
		bookDetails.Autore = updateBook.Autore
	}
	if updateBook.Pubblicazione != "" {
		bookDetails.Pubblicazione = updateBook.Pubblicazione
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stringId := params["bookid"]
	id, err := strconv.ParseInt(stringId, 0, 0)
	if err != nil {
		return
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
