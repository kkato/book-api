# Book API

書籍を管理するシンプルなREST APIです。Go標準ライブラリのみで実装されています。

## 機能

- 書籍のCRUD操作
- インメモリストレージ
- 外部依存なし

## 起動方法

```bash
# サーバーを起動
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

## APIエンドポイント

| メソッド | エンドポイント | 説明 |
|--------|----------|-------------|
| GET | `/books` | すべての書籍を取得 |
| POST | `/books` | 新しい書籍を作成 |
| GET | `/books/{id}` | IDで書籍を取得 |
| PUT | `/books/{id}` | 書籍を更新 |
| DELETE | `/books/{id}` | 書籍を削除 |

## 使用例

```bash
# 書籍を作成
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "title": "プログラミング言語Go",
    "author": "Alan Donovan",
    "isbn": "978-0134190440",
    "published_at": "2015-10-26T00:00:00Z"
  }'

# すべての書籍を取得
curl http://localhost:8080/books

# 特定の書籍を取得
curl http://localhost:8080/books/1

# 書籍を更新
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "プログラミング言語Go（改訂版）",
    "author": "Alan Donovan & Brian Kernighan",
    "isbn": "978-0134190440",
    "published_at": "2015-10-26T00:00:00Z"
  }'

# 書籍を削除
curl -X DELETE http://localhost:8080/books/1
```
