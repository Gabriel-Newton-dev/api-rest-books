package routes

import (
	"github.com/Gabriel-Newton-dev/api-rest-books/book"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/books", book.DisplaysAllBooks)
	r.Run()

}
