package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisClient = rdb
}

func CacheMessageHistory(senderID, receiverID int, messages []byte) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("message_history_%d_%d", senderID, receiverID)
	err := redisClient.Set(ctx, cacheKey, messages, 0).Err()
	if err != nil {
		log.Printf("Failed to cache message history in Redis: %s\n", err)
	}
}
