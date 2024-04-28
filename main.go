package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

type Book struct {
	ID 			string	`json:"id"`
	Title 		string	`json:"title"`
	Author 		string	`json:"author"`
	Quantity 	int		`json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel", Quantity: 4},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 2},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}