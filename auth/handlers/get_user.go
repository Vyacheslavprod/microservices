package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(ctx *gin.Context) {
	userId, err := utils.ExtractUserID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен: " + err.Error()})
		return
	}
	var user models.User
	result := database.DB.Where("ID = ?", userId).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}
	// Создаем анонимную структуру с только id и email
	userResponse := struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Email: user.Email,
	}
	ctx.JSON(http.StatusOK, gin.H{"user": userResponse})
}