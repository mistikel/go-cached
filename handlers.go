package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Index
func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hi, welcome to go-cached")
}

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	books, err := Gets()
	HandleError(err)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		panic(err)
	}
}

// Get a book
func GetBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(p.ByName("id"))
	books, err := Get(id)
	HandleError(err)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		panic(err)
	}
}

// Store a book
func StoreBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hi, welcome to go-cached")
}

// Update a book
func UpdateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hi, welcome to go-cached")
}

func DeleteBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	id, _ := strconv.Atoi(p.ByName("id"))
	err := Delete(id)
	if err != nil {
		w.Write([]byte("No book removed"))
	}
	w.Write([]byte("Book was removed "))
}
