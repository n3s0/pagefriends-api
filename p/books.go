package p

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

/*
Data intended for testing outside of a database.
*/
var books = []Book{
	{ID: "1", Title: "Mistborn", Author: "Brandon Sanderson"},
	{ID: "2", Title: "Iron Flame", Author: "Rebecca Yarros"},
	{ID: "3", Title: "Onyx Storm", Author: "Rebecca Yarros"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
