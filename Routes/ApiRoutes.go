package Routes

import (
	"books-api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", Controllers.GetBooks)
	router.GET("/books/:id", Controllers.GetBookByID)
	router.POST("/books", Controllers.PostBooks)

	return router
}
