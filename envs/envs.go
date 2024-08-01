package envs

import (
	"os"
	"github.com/joho/godotenv"
)

// Хранение данных значений ENV
var ServerEnvs Envs

// Структура для хранения значений ENV
type Envs struct {
	MONGO_INITDB_ROOT_PASSWORD string
	MONGO_INITDB_ROOT_USERNAME string
	MONGO_INITDB_PORT          string
	MONGO_INITDB_HOST          string
	NOTES_PORT                 string
}

// Инициализация значений ENV
func LoadEnvs() error {
	// Загрузка значений ENV.
	// Если файл .env не найден, то выводим сообщение
	if err := godotenv.Load(); err != nil {
		return err
	}
	// Инициализация значений ENV
	ServerEnvs.NOTES_PORT = os.Getenv("NOTES_PORT")
	ServerEnvs.MONGO_INITDB_ROOT_USERNAME = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	ServerEnvs.MONGO_INITDB_ROOT_PASSWORD = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	ServerEnvs.MONGO_INITDB_PORT = os.Getenv("MONGO_INITDB_PORT")
	ServerEnvs.MONGO_INITDB_HOST = os.Getenv("MONGO_INITDB_HOST")
	return nil
}