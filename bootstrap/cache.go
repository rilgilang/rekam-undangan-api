package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"github.com/rilgilang/rekam-undangan-api/config/dotenv"
)

func NewCache(config *dotenv.Config) *redis.Client {
	// Redis
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
}
