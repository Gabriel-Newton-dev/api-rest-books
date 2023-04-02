package routes

import (
	"github.com/Gabriel-Newton-dev/API_Rest_Golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/books", controllers.DisplaysAllBooks)
	r.Run()
}
