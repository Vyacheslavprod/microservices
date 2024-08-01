package database

import (
	"context"
	"fmt"
	"log"
	"github.com/vyacheslavprod/microservices/envs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Объявление переменной MongoClient, хранящей ссылку на экземпляр клиента MongoDB
var MongoClient *mongo.Client

func InitDatabase() error {
	env := &envs.ServerEnvs
	// формирование URI для подключения к MongoDB.
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.MONGO_INITDB_ROOT_USERNAME, env.MONGO_INITDB_ROOT_PASSWORD, env.MONGO_INITDB_HOST, env.MONGO_INITDB_PORT)
	log.Println("URI: " + mongoURI)
	// Создаем новый контекст с таймаутом и предусматриваем его корректное завершение.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Отмена контекста при завершении работы функции
	defer cancel()
	// Создание клиента MongoDB и попытка подключения к серверу
	mongo, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	// Проверка на наличие ошибки при попытке подключения.
	if err != nil {
		return err
	} else {
		// Сохранение клиента в глобальной переменной
		MongoClient = mongo
	}
	// Проверка соединения
	mongoErr := MongoClient.Ping(ctx, readpref.Primary())
	if mongoErr != nil {
		return mongoErr
	}
	return nil
}