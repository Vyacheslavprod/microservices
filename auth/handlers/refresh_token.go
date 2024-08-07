package handlers

import (
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func refreshToken(ctx *gin.Context) {
	var token RefreshTokenRequest

	// Извлечение данных из тела запроса и заполнение структуры
	if err := ctx.ShouldBindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при разборе данных запроса"})
		return
	}

	userId, err := utils.ValidateRefreshToken(token.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный refresh токен" + err.Error()})
		return
	}

	tokens, err := utils.GenerateTokens(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токены"})
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}