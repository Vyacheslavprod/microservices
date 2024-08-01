package database

import (
	"fmt"
	"github.com/vyacheslavprod/microservices/envs"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

// Подключение к Redis
func InitRedis() error {
	// формирование URI для подключения к Redis
	redisUri := fmt.Sprintf("%s:%s", envs.ServerEnvs.REDIS_HOST, envs.ServerEnvs.REDIS_PORT)

	// Создание нового клиента Redis и присваивание его глобальной
	// переменной [database.RedisClient]
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisUri,
		Password: "", // Пароль не нужен
		DB:       0,  // База данных 0 (по умолчанию)
	})
	// Проверка соединения, отправка команды PING
	// в ответ при успешном подключении к Redis нам
	// придет "PONG" или ошибка
	status := RedisClient.Ping()
	if status.Val() == "PONG" {
		return nil
	} else {
		return fmt.Errorf("ошибка при подключении к Redis: %v", status)
	}

}