package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kkato/book-api/internal/model"
)

var ErrBookNotFound = errors.New("book not found")

type BookRepository interface {
	Create(ctx context.Context, book model.Book) (model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, id int) (model.Book, error)
	Update(ctx context.Context, id int, book model.Book) (model.Book, error)
	Delete(ctx context.Context, id int) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAll(ctx context.Context) ([]model.Book, error) {
	query := `SELECT id, title, author, isbn, published_at, created_at, updated_at FROM books ORDER BY id`

	sqlRows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer sqlRows.Close()

	books := make([]model.Book, 0)
	for sqlRows.Next() {
		var book model.Book
		if err := sqlRows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.PublishedAt, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, sqlRows.Err()
}

func (r *bookRepository) Create(ctx context.Context, book model.Book) (model.Book, error) {
	query := `INSERT INTO books (title, author, isbn, published_at) VALUES (?, ?, ?, ?)`

	res, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.ISBN, book.PublishedAt)
	if err != nil {
		return model.Book{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Book{}, err
	}

	return r.GetByID(ctx, int(id))
}

func (r *bookRepository) GetByID(ctx context.Context, id int) (model.Book, error) {
	query := `SELECT id, title, author, isbn, published_at, created_at, updated_at FROM books WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)

	var book model.Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.PublishedAt, &book.CreatedAt, &book.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Book{}, ErrBookNotFound
	}
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (r *bookRepository) Update(ctx context.Context, id int, book model.Book) (model.Book, error) {
	query := `UPDATE books SET title = ?, author = ?, isbn = ?, published_at = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`

	res, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.ISBN, book.PublishedAt, id)
	if err != nil {
		return model.Book{}, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return model.Book{}, err
	}
	if rows == 0 {
		return model.Book{}, ErrBookNotFound
	}

	return r.GetByID(ctx, id)
}

func (r *bookRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM books WHERE id = ?`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrBookNotFound
	}

	return nil
}
