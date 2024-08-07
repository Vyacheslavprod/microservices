package utils

import (
	"auth/envs"
	"auth/models"
	"fmt"
	"time"
	"strings"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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

// HashPassword хеширует пароль с использованием bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 10 - это стоимость хеширования
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash проверяет хеш пароля с использованием bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Функция для генерации JWT токена
func GenerateTokens(userID uint) (models.Tokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Срок действия токена 24 часа
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 600).Unix(), // Срок действия месяц
	})
	signedAccessToken, _ := accessToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))

	signedRefreshToken, _ := refreshToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))

	return models.Tokens{AccessToken: signedAccessToken, RefreshToken: signedRefreshToken}, nil
}

// validateRefreshToken валидирует JWT токен
func ValidateRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDValue, ok := claims["user_id"].(float64) // Приведение к float64
		if !ok {
			return 0, fmt.Errorf("user_id claim is not a float64")
		}
		return uint(userIDValue), nil // Конвертация float64 в uint
	} else {
		return 0, fmt.Errorf("недействительный токен")
	}
}