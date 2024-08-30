package p

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/authors", getAuthors)

	router.Run("localhost:8080")
}
