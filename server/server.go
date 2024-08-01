package server

import (
	"log"
	"github.com/vyacheslavprod/microservices/database"
	"github.com/vyacheslavprod/microservices/envs"
)

func InitServer() {
	// Инициализация внешних значений ENV
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		// Вывод сообщения об ошибке
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}
	// Инициализация базы данных
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
	}
	errRedis := database.InitRedis()
	if errRedis != nil {
		log.Fatal("Ошибка подключения к Redis: ", errRedis)
	} else {
		log.Println("Успешное подключение к Redis")
	}
}

func StartServer() {
	// Инициализация роутеров
	// Запуск сервера
	InitRotes()
}