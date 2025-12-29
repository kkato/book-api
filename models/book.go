package models

import (
	"errors"
	"sync"
	"time"
)

type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookStore struct {
	mu    sync.RWMutex
	books map[string]*Book
}

func NewBookStore() *BookStore {
	return &BookStore{
		books: make(map[string]*Book),
	}
}

func (s *BookStore) Create(book *Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[book.ID]; exists {
		return errors.New("book with this ID already exists")
	}

	now := time.Now()
	book.CreatedAt = now
	book.UpdatedAt = now
	s.books[book.ID] = book
	return nil
}

func (s *BookStore) GetAll() []*Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	books := make([]*Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	return books
}

func (s *BookStore) GetByID(id string) (*Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, exists := s.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (s *BookStore) Update(id string, book *Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}

	book.ID = id
	book.UpdatedAt = time.Now()
	if s.books[id].CreatedAt.IsZero() {
		book.CreatedAt = time.Now()
	} else {
		book.CreatedAt = s.books[id].CreatedAt
	}
	s.books[id] = book
	return nil
}

func (s *BookStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}

	delete(s.books, id)
	return nil
}
