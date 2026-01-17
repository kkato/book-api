package main

import (
	"log"
	"net/http"

	"github.com/kkato/book-api/handlers"
	"github.com/kkato/book-api/models"
)

func main() {
	store := models.NewBookStore()
	bookHandler := handlers.NewBookHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", bookHandler.ListBooks)
	mux.HandleFunc("POST /books", bookHandler.CreateBook)
	mux.HandleFunc("GET /books/{id}", bookHandler.GetBook)
	mux.HandleFunc("PUT /books/{id}", bookHandler.UpdateBook)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
