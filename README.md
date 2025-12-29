# Book API

A simple REST API for managing books, built with Go's standard library.

## Features

- CRUD operations for books
- In-memory storage
- No external dependencies

## Getting Started

```bash
# Run the server
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/books` | List all books |
| POST | `/books` | Create a new book |
| GET | `/books/{id}` | Get a book by ID |
| PUT | `/books/{id}` | Update a book |
| DELETE | `/books/{id}` | Delete a book |

## Example

```bash
# Create a book
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "title": "The Go Programming Language",
    "author": "Alan Donovan",
    "isbn": "978-0134190440",
    "published_at": "2015-10-26T00:00:00Z"
  }'

# Get all books
curl http://localhost:8080/books

# Get a specific book
curl http://localhost:8080/books/1

# Update a book
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language (Updated)",
    "author": "Alan Donovan & Brian Kernighan",
    "isbn": "978-0134190440",
    "published_at": "2015-10-26T00:00:00Z"
  }'

# Delete a book
curl -X DELETE http://localhost:8080/books/1
```
