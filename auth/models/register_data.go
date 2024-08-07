package models

/// Структура данных хранения данных для регистрации пользователя
type RegisterData struct {
	// Требуется валидный email
	Email string `json:"email" binding:"required,email"`
	// Требуется валидный пароль с минимальной длиной 8 символов
	Password string `json:"password" binding:"required,min=8"`
}