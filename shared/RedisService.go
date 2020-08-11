package shared

import "github.com/go-redis/redis/v8"

type RedisService struct {
	RedisClient *redis.Client
}

func NewRedisService() *RedisService {
	return &RedisService{RedisClient: redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})}
}
