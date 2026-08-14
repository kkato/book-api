package main

import (
	"log"
	"net/http"

	"github.com/kkato/book-api/internal/database"
	"github.com/kkato/book-api/internal/handler"
	"github.com/kkato/book-api/internal/repository"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewBookRepository(db)
	bookHandler := handler.NewBookHandler(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", bookHandler.GetAll)
	mux.HandleFunc("POST /books", bookHandler.Create)
	mux.HandleFunc("GET /books/{id}", bookHandler.GetByID)
	mux.HandleFunc("PUT /books/{id}", bookHandler.Update)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.Delete)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
