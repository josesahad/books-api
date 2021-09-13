package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// book represents a book.
type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var books = []book{
	{ID: "1", Title: "Ulysses", Author: "James Joyce", Year: 1920},
	{ID: "2", Title: "The Aleph", Author: "Jorge Luis Borges", Year: 1949},
	{ID: "3", Title: "Don Quixote", Author: "Miguel de Cervantes", Year: 1605},
}

// getBooks responds with the list of books as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks creates a new book from JSON received in the request body.
func postBooks(c *gin.Context) {
	var bookToAdd book

	if err := c.BindJSON(&bookToAdd); err != nil {
		return
	}

	books = append(books, bookToAdd)
	c.IndentedJSON(http.StatusCreated, bookToAdd)
}

// getBookByID retrieves a book by id.
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	router.Run()
}
