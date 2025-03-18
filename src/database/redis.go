package database

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func ConnectRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatalln("REDIS_URL environment variable is not set")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // No password set
		DB:       0,  // Use default DB
	})

	if err := client.Ping().Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	Redis = client
}
