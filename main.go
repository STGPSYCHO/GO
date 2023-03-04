package main

import (
	"net/http"

	"github.com/STGPSYCHO/GO/controllers"
	"github.com/STGPSYCHO/GO/models"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	models.ConnectDB()

	route.GET("/users", controllers.GetAllUsers)
	route.POST("/users", controllers.CreateUser)
	route.GET("/users/:id", controllers.GetUser)
	route.PATCH("/users/:id", controllers.UpdateUser)
	route.DELETE("/users/:id", controllers.DeleteUser)

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	route.Run()
}
