package main

import (
	"net/http"

	"github.com/STGPSYCHO/GO/models"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	models.ConnectDB()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	route.Run()
}
