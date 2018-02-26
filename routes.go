package main

import (
	"github.com/julienschmidt/httprouter"
)

func Routes(r *httprouter.Router) {
	r.GET("/", Index)
	r.GET("/books", GetBooks)
	r.GET("/books/:id", GetBook)
	r.POST("/books", StoreBook)
	r.PUT("/books/:id", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)
}
