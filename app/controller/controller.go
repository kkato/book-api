package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
	router.PATCH("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	return router
}
