package Controllers

import (
	"books-api/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var books = []Models.Book{
	{ID: "1", Title: "Ulysses", Author: "James Joyce", Year: 1920},
	{ID: "2", Title: "The Aleph", Author: "Jorge Luis Borges", Year: 1949},
	{ID: "3", Title: "Don Quixote", Author: "Miguel de Cervantes", Year: 1605},
}

// getBooks responds with the list of books as JSON.
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks creates a new book from JSON received in the request body.
func PostBooks(c *gin.Context) {
	var bookToAdd Models.Book

	if err := c.BindJSON(&bookToAdd); err != nil {
		return
	}

	books = append(books, bookToAdd)
	c.IndentedJSON(http.StatusCreated, bookToAdd)
}

// getBookByID retrieves a book by id.
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
