# book-api

An API for managing books.

## Endpoints
- `POST /books` -> Create a new book
- `GET /books` -> Return a list of all books
- `GET /books/{id}` -> Return details of a specific book
- `PATCH /books/{id}` -> Update a specific book
- `DELETE /books/{id}` -> Delete a specific book
- All responses are returned in JSON format.
- The HTTP response status code for successful requests is 200, and requests to non-existent endpoints return 404.

## Endpoint Details

### `POST /books`Endpoint
Create a new book.

- Request format:
  - Required fields: `title`, `author`, `publication_year`, `genre`, `price`

- Success response:
```json
{
    "message": "Book successfully created!",
    "book": {
        "id": 1,
        "title": "Go Programming",
        "author": "John Doe",
        "publication_year": "2022",
        "genre": "Programming",
        "price": "3500",
        "created_at": "2024-09-07T10:00:00Z",
        "updated_at": "2024-09-07T10:00:00Z"
    }
}
```

- Failure response:
```json
{
    "message": "Book creation failed",
    "required": "title, author, publication_year, genre, price"
}
```

### `GET /books`Endpoint

Returns a list of all books.

- Request format: `GET /books/`
- Response format:
```json
{
    "books": [
        {
            "id": 1,
            "title": "Go Programming",
            "author": "John Doe",
            "publication_year": "2022",
            "genre": "Programming",
            "price": "3500"
        },
        {
            "id": 2,
            "title": "Introduction to Kubernetes",
            "author": "Jane Smith",
            "publication_year": "2021",
            "genre": "Technology",
            "price": "4000"
        }
    ]
}
```


### `GET /books/{id}`Endpoint

Returns the details of the book with the specified id.

- Request format: `GET /books/{id}`
- Response format:
```json
{
    "message": "Book details by id",
    "book": {
        "id": 1,
        "title": "Go Programming",
        "author": "John Doe",
        "publication_year": "2022",
        "genre": "Programming",
        "price": "3500"
    }
}
```

### `PATCH /books/{id}`Endpoint

Updates the book with the specified id and returns the updated book.

- Request format: PATCH /books/{id}
  - Fields: One or more of `title`, `author`, `publication_year`, `genre`, `price`
- Success response:
```json
{
    "message": "Book successfully updated!",
    "book": {
        "id": 1,
        "title": "Go Programming Advanced",
        "author": "John Doe",
        "publication_year": "2022",
        "genre": "Programming",
        "price": "4000"
    }
}
```

- Failure response:
```json
{
    "message": "Book update failed",
    "required": "title, author, publication_year, genre, price"
}
```

### `DELETE /books/{id}`Endpoint

Deletes the book with the specified id.

- Request format: `DELETE /books/{id}`
- Success response:
```json
{
    "message": "Book successfully removed!"
}
```

- Failure response:
```json
{
    "message": "No book found"
}
```
