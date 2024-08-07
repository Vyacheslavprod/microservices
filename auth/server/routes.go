package server

import (
	"auth/envs"
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func InitRotes() {
	// Инициализация  роута (по умолчанию)
	router := gin.Default()
	// Создание пользователя
	router.PUT("/user", handlers.RegisterUserHandler)
	// Авторизация пользователя
	router.POST("/user", handlers.SignInHandler)
	// Обновление токена
	router.POST("/refresh", handlers.RefreshTokenHandler)

	// Перехватчик для авторизованных пользователей
	auth := router.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		// Получение данных пользователя, если пропустит перехватчик
		router.GET("/user", handlers.GetUserHandler)
	}

	router.Run(":" + envs.ServerEnvs.AUTH_PORT)

}