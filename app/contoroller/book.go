package controller

import (
	"net/http"

	"github.com/kkato/book-api/app/model"

	"github.com/gin-gonic/gin"
)

var err error

func getBooks(c *gin.Context) {
	books := model.GetBooks()
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func getBook(c *gin.Context) {
	book := model.GetBook(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func createBook(c *gin.Context) {
	var book model.Book
	err = c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Book creation failed", "required": "title, author, publication_year, genre, price"})
		return
	}
	err = model.CreateBook(book)
	createdBook := model.GetBook(book.id)
	c.JSON(http.StatusOK, gin.H{"message": "Recipe successfully created!", "recipe": createdBook})
}

func updateBook(c *gin.Context) {
	book := model.UpdateBook(c.Param("id"), c.PostForm("title"), c.PostForm("author"))
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func deleteBook(c *gin.Context) {
	book := model.DeleteBook(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"book": book})
}
