package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"errors"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel", Quantity: 4},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 2},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 1},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*Book, error) {
	for i, b := range books {
		if (b.ID == id) {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not foud")
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}
