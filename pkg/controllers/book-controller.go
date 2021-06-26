package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KaT0819/sample-go-gorilla-api/pkg/models"
	"github.com/KaT0819/sample-go-gorilla-api/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)

	b := Book.Create()

	// response
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAll()

	// response
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("パラメータ不正")
	}

	book, _ := models.Get(id)

	// response
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	param := &models.Book{}
	utils.ParseBody(r, param)

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("パラメータ不正")
	}

	book, db := models.Get(id)

	if param.Name != "" {
		book.Name = param.Name
	}
	if param.Author != "" {
		book.Author = param.Author
	}
	if param.Publication != "" {
		book.Publication = param.Publication
	}

	db.Save(&book)

	// response
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("パラメータ不正")
	}

	book := models.Delete(id)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
