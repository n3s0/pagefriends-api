package p

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/api/v1/books/", getBooks)
	router.GET("/api/v1/authors/", getAuthors)

	router.Run("localhost:8080")
}
