package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Author struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	MiddleInitial string `json:"middle_initial"`
	LastName      string `json:"last_name"`
}

func CreateAuthor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "created author"})
}

func GetAuthors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, authors)
}

func GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": id})
}

func UpdateAuthor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updated author"})
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deleted author: " + id})
}
