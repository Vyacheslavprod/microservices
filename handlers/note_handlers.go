package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Обработка запроса для получения заметки по ID
func GetNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetNoteHandler")
}

// Обработка запроса для получения всех заметок
func GetNotesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetNotesHandler")
}

// Обработка запроса для удаления заметки по ID
func DeleteNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "DeleteNoteHandler")
}

// Обработка запроса для редактирования заметки по ID
func UpdateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "UpdateNoteHandler")
}

// Обработка запроса для создания заметки
func CreateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "CreateNoteHandler")
}