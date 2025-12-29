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

	mux.HandleFunc("GET /books", bookHandler.HandleBooks)
	mux.HandleFunc("POST /books", bookHandler.HandleBooks)
	mux.HandleFunc("GET /books/{id}", bookHandler.HandleBook)
	mux.HandleFunc("PUT /books/{id}", bookHandler.HandleBook)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.HandleBook)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
