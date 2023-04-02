package book

import (
	"github.com/gin-gonic/gin"
)

func DisplaysAllBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"book": "Os 3 porquinhos"})
}
