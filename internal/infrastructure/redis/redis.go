package redis

import (
	"context"
	"learning-keycloak/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	RDS *redis.Client
}

var redisClient *RedisClient

func NewConnectToRedis(ctx context.Context) *RedisClient {

	config := config.NewRedisConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST + ":" + config.REDIS_PORT,
		Password: config.REDIS_PASSWORD,
		DB:       config.REDIS_DATABASES,
	})

	_, errConnect := client.Ping(ctx).Result()

	if errConnect != nil {
		panic("Failed To Connect Redis!: " + errConnect.Error())
	}

	if redisClient == nil {
		redisClient = &RedisClient{
			RDS: client,
		}
		log.Println("Success Connect To Redis ✅🎌")
	}

	return redisClient
}

func NewGetInstaceRedis() *RedisClient {
	if redisClient == nil {
		panic("Redis Client is not initialized. Please call NewConnectToRedis first.")
	}
	return redisClient
}
