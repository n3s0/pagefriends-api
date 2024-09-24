package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
		/* Routes for Books */
		v1.GET("books", GetBooks)
		v1.GET("books/:id", GetBookById)
		v1.POST("books", CreateBook)
		v1.PUT("books/:id", UpdateBook)
		v1.DELETE("books/:id", DeleteBook)
		/* Routes for Authors */
		v1.GET("authors", GetAuthors)
		v1.GET("authors/:id", GetAuthorById)
		v1.POST("authors", CreateAuthor)
		v1.PUT("authors/:id", UpdateAuthor)
		v1.DELETE("authors/:id", DeleteAuthor)
	}

	r.Run("localhost:8080")
}
