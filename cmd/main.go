package main

import "github.com/gin-gonic/gin"

func main() {
	// Создайте новый роутер по умолчанию у gin
	router := gin.Default()
	// Регистрируем обработчик GET-запроса /ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	// Запускаем роутер
	router.Run(":9100")
}