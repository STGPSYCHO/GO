package controllers

import (
	"net/http"

	"strings"

	"github.com/STGPSYCHO/GO/models"
	"github.com/gin-gonic/gin"
)

type AddProductInput struct {
	Product  string `json:"product"`
	Quantity string `json:"quantity"`
}

// GET /products
// Получаем список всех продуктов
func GetAllProducts(context *gin.Context) {

	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Нет подходящих записей"})
		return
	}

	context.HTML(
		http.StatusOK,
		"products.html",
		gin.H{"products": products},
	)
}

// POST /add-cart
// Добавляем корзину в куку
func AddProductToCart(context *gin.Context) {

	var products models.Product
	var input AddProductInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", input.Product).First(&products).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Запись не существует"})
		return
	}

	var cart []string
	cookie, err := context.Cookie("cart")

	if cookie != "" {
		cart = strings.Split(cookie, ",")
		cart = append(cart, input.Product+":"+input.Quantity)
		cookie = strings.Join(cart, ",")
	} else {
		cookie = input.Product + ":" + input.Quantity
	}
	if err != nil {
	}
	context.SetCookie("cart", cookie, 3600, "/", context.Request.URL.Hostname(), false, true)
	context.JSON(http.StatusOK, gin.H{"Успех": cookie})
	// context.Redirect(http.StatusMovedPermanently, "/products")
}

// POST /remove-cart
// Удаляем из корзины куку
func RemoveProductToCart(context *gin.Context) {

	var products models.Product
	var input AddProductInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", input.Product).First(&products).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Запись не существует"})
		return
	}

	cookie, err := context.Cookie("cart")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Нет таких товаров в корзине"})
		return
	}
	if cookie != "" {
		cart := strings.Split(cookie, ",")
		idx := Find(cart, input.Product+":"+input.Quantity)
		cart = remove(cart, idx)
		cookie = strings.Join(cart, ",")
	} else {
		cookie = input.Product + ":" + input.Quantity
	}
	context.SetCookie("cart", cookie, 3600, "/", context.Request.URL.Hostname(), false, true)
	// context.Redirect(http.StatusMovedPermanently, "/products")
}

// GET /cookie
// Удаляем из корзины куку
func GetCookie(context *gin.Context) {

	cookie, err := context.Cookie("cart")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"cookie": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"cookie": cookie})
}

// GET /cart
// Получаем корзину товаров
func GetCart(context *gin.Context) {

	var arr []AddProductInput

	cookie, err := context.Cookie("cart")
	cart := strings.Split(cookie, ",")
	strings.Cut()
	for index, value := range cart {
		if index != 0 {
		}
		b, a, f := strings.Cut(value, ":")
		if f == true {
			input := AddProductInput{Product: b, Quantity: a}
			arr = append(arr, input)
		}
	}
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"cookie": "Ошибка получения корзины"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"cookie": cookie})
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
