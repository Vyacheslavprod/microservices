package server

import (
	"github.com/vyacheslavprod/microservices/notes/envs"
	"github.com/vyacheslavprod/microservices/notes/handlers"

	"github.com/gin-gonic/gin"
)

func InitRotes() {
	// Инициализация  роута (по умолчанию)
	router := gin.Default()

	auth := router.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		// Создание заметки
		router.PUT("/note", handlers.CreateNoteHandler)
		// Удаление заметки
		router.DELETE("/note/:id", handlers.DeleteNoteHandler)
		// Получение заметки
		router.GET("/note/:id", handlers.GetNoteHandler)
		// Редактирование заметки
		router.POST("/note/:id", handlers.UpdateNoteHandler)
		// Получение списка всех заметок
		router.GET("/notes", handlers.GetNotesHandler)
	}

	// Запуск сервера на порту 9100
	router.Run(":" + envs.ServerEnvs.NOTES_PORT)
}
