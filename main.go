package main

import (
	"net/http"

	"github.com/STGPSYCHO/GO/controllers"
	"github.com/STGPSYCHO/GO/models"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.LoadHTMLGlob("templates/*")

	models.ConnectDB()

	// Работа с пользователем
	route.GET("/users", controllers.GetAllUsers)
	route.GET("/users-by-headers", controllers.GetAllUsersByHeaders)
	route.POST("/users", controllers.CreateUser)
	route.GET("/users/:id", controllers.GetUser)
	route.PATCH("/users/:id", controllers.UpdateUser)
	route.DELETE("/users/:id", controllers.DeleteUser)

	// Работа с продуктами
	route.GET("/products", controllers.GetAllProducts)
	route.POST("/add-cart", controllers.AddProductToCart)
	route.POST("/remove-cart", controllers.RemoveProductToCart)
	route.GET("/cookie", controllers.GetCookie)
	route.GET("/cart", controllers.GetCart)

	// Базовый эндпоинт
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	route.Run()
}
