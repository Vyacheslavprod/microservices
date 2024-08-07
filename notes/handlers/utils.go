package handlers

import (
	"fmt"
	"github.com/vyacheslavprod/microservices/notes/envs"
	"strings"

	"github.com/golang-jwt/jwt"
)

// extractUserID извлекает ID пользователя из JWT токена
func ExtractUserID(tokenString string) (uint, error) {
	// Отсечение префикса "Bearer " из заголовка
	str := strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer"))

	// Проверяем, что токен валиден
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		// Убедитесь, что алгоритм подписи, который вы ожидаете, соответствует 'jwt.SigningMethodHS256'
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный алгоритм подписи: %v", token.Header["alg"])
		}

		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"]

		if userIDFloat, ok := userID.(float64); ok {
			return uint(userIDFloat), nil // Преобразуем float64 в uint
		}
	}

	return 0, fmt.Errorf("невозможно извлечь user_id из токена")
}