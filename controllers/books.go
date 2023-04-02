package controlles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplaysAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ID":   "1",
		"Book": "Alice no pais das maravilhas."})
}
