package controllers

import (
	"net/http"

	"github.com/STGPSYCHO/GO/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Email      string `json:"email" binding:"required"`
}

type UpdateUserInput struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

// GET /users-by-headers
// Получаем список всех юзеров
func GetAllUsersByHeaders(context *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	if (context.ContentType() == "application/json") || (context.ContentType() == "text/plain") {
		switch context.GetHeader("Accept") {
		case "application/json":
			context.JSON(http.StatusOK, gin.H{"users": users})
		case "text/plain":

			// Заменить Success на свой вывод юзеров.
			// context.Writer.WriteHeader(http.StatusOK)
			// context.Writer.Header().Set("Content-Type", "text/plain")
			// context.Writer.Write([]byte("Success"))
		default:
			context.JSON(http.StatusNotFound, gin.H{"error": "Страница не найдена"})
		}
	} else {
		context.JSON(http.StatusNotFound, gin.H{"error": "Страница не найдена"})
	}

}

// GET /users
// Получаем список всех юзеров
func GetAllUsers(context *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	if q, ok := context.GetQuery("format"); ok {
		if q == "json" {
			context.JSON(http.StatusOK, gin.H{"users": users})
		}
		if q == "text" {
			// Заменить Success на свой вывод юзеров.
			// context.Writer.WriteHeader(http.StatusOK)
			// context.Writer.Header().Set("Content-Type", "text/plain")
			// context.Writer.Write([]byte("Success"))
		}
	} else {
		context.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"users": users},
		)
	}
}

// GET /users/:id
// Получаем юзера по его id
func GetUser(context *gin.Context) {
	var users models.User

	if err := models.DB.Where("id = ?", context.Param("id")).First(&users).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Запись не существует"})
		return
	}

	// context.JSON(http.StatusOK, gin.H{"users": users})
	context.HTML(
		http.StatusOK,
		"user.html",
		gin.H{
			"ID":         users.ID,
			"First_name": users.First_name,
			"Last_name":  users.Last_name,
			"Email":      users.Email,
		},
	)
}

// POST /users
// Создание юзера
func CreateUser(context *gin.Context) {
	var input CreateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{First_name: input.First_name, Last_name: input.Last_name, Email: input.Email}
	models.DB.Create(&user)

	context.JSON(http.StatusCreated, gin.H{"users": user})
}

// PATCH /users/:id
// Изменения информации
func UpdateUser(context *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её менять
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatesUser := models.User{First_name: input.First_name, Last_name: input.Last_name, Email: input.Email}

	models.DB.Model(&user).Updates(&updatesUser)

	context.JSON(http.StatusOK, gin.H{"users": user})
}

// DELETE /users/:id
// Удаление
func DeleteUser(context *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её удалять
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&user)

	context.JSON(int(http.StatusNoContent), gin.H{"users": true})
}
