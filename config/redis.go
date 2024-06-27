package config

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func SetupRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis server address
        Password: "",                // No password set
        DB:       0,                 // Use default DB
    })

    // Ping the Redis server to check connectivity
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
    log.Println("Connected to Redis!")
}
