package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/kkato/book-api/config"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title" binding:"required"`
	Author          string    `json:"author" binding:"required"`
	PublicationYear int       `json:"publication_year" binding:"required"`
	Genre           string    `json:"genre" binding:"required"`
	Price           int       `json:"price" binding:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

var Db *sql.DB
var err error

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	dropTableSQ := "DROP TABLE IF EXISTS books;"
	_, err = Db.Exec(dropTableSQ)
	if err != nil {
		log.Fatalln(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title STRING NOT NULL,
		author STRING NOT NULL,
		publication_year INTEGER NOT NULL,
		genre STRING NOT NULL,
		price INTEGER NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = Db.Exec(createTableSQL)

	if err != nil {
		log.Fatalln(err)
	}

	insertSQL1 := `INSERT INTO books (title, author, publication_year, genre, price)
		VALUES ('Go Programming', 'John Doe', '2022', 'Programming', 3500)`
	_, err = Db.Exec(insertSQL1)
	if err != nil {
		log.Fatalln(err)
	}

	insertSQL2 := `INSERT INTO books (title, author, publication_year, genre, price)
		VALUES ('Introduction to Kubernetes', 'Jane Smith', '2021', 'Technology', 4000)`
	_, err = Db.Exec(insertSQL2)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetBooks() (books []Book) {
	cmd := "SELECT * FROM books"
	rows, _ := Db.Query(cmd)
	defer rows.Close()

	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Price, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		books = append(books, book)
	}
	return books
}

func GetBook(id string) (book Book) {
	cmd := "SELECT * FROM books WHERE id = ?"
	rows := Db.QueryRow(cmd, id)
	err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublicationYear, &book.Genre, &book.Price, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	return book
}

func CreateBook(title string, author string, publicationYear int, genre string, price int) (id string) {
	cmd := "INSERT INTO books (title, author, publication_year, genre, price, created_at, updated_at) VALUES (?, ?, ?, ?, ?, datetime('now'), datetime('now'))"
	_, err = Db.Exec(cmd, title, author, publicationYear, genre, price)
	if err != nil {
		log.Fatalln(err)
	}
	return id
}

func UpdateBook(id int, title string, author string, publicationYear int, genre string, price int) (err error) {
	cmd := "UPDATE books SET title = ?, author = ?, publication_year = ?, genre = ?, price = ?, updated_at = datetime('now') WHERE id = ?"
	_, err = Db.Exec(cmd, title, author, publicationYear, genre, price, id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func DeleteBook(id string) (err error) {
	cmd := "DELETE FROM books WHERE id = ?"
	_, err = Db.Exec(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
