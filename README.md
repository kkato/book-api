# book-api
書籍管理のためのHTTPエンドポイントを作成する。

## エンドポイント
- `POST /books` -> 書籍を作成
- `GET /books` -> 全書籍の一覧を返す
- `GET /books/{id}` -> 指定した書籍を1つ返す
- `PATCH /books/{id}` -> 指定した書籍を更新
- `DELETE /books/{id}` -> 指定した書籍を削除
- レスポンスはすべてJSON形式で返す。
- エンドポイントに対するHTTPレスポンスステータスコードは、成功時に`200`を返し、存在しないエンドポイントへのリクエストには`404`を返す。

## エンドポイント詳細

### `POST /books`エンドポイント
書籍を新規作成する。

- リクエスト形式
  - 必須フィールド: `title`, `author`, `publication_year`, `genre`, `price`

- 成功時のレスポンス:
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

失敗時のレスポンス:
```json
{
    "message": "Book creation failed",
    "required": "title, author, publication_year, genre, price"
}
```

### `GET /books`エンドポイント

すべての書籍を返す。

- リクエスト形式: `GET /books/`
- レスポンス形式
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


### `GET /books/{id}`エンドポイント

指定したidの書籍を返す。

- リクエスト形式: `GET /books/{id}`
- レスポンス形式:
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

### `PATCH /books/{id}`エンドポイント

指定したidの書籍を更新し、更新された書籍を返す。

- リクエスト形式: PATCH /books/{id}
  - フィールド: `title`, `author`, `publication_year`, `genre`, `price` のいずれか
- 成功時のレスポンス:
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

- 失敗時のレスポンス:
```json
{
    "message": "Book update failed",
    "required": "title, author, publication_year, genre, price"
}
```

### `DELETE /books/{id}`エンドポイント

指定したidの書籍を削除する。

- リクエスト形式: `DELETE /books/{id}`
- 成功時のレスポンス:
```json
{
    "message": "Book successfully removed!"
}
```

- 失敗時のレスポンス:
```json
{
    "message": "No book found"
}
```






