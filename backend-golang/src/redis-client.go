package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisClient interface {
	Incr() int
	Close()
}

type RedisClientImpl struct {
	client *redis.Client
	ctx    context.Context
	key    string
}

func (r RedisClientImpl) Incr() int {
	res, err := r.client.Incr(r.ctx, r.key).Result()
	if err != nil {
		log.Println("Redis call failed")
		log.Fatal(err)
	}
	return int(res)
}

func (r RedisClientImpl) Close() {
	_ = r.client.Close()
}

// NewRedisClient Factory function
func NewRedisClient(host string, key string) RedisClient {
	address := host + ":6379"
	log.Printf("Connecting to Redis at %s ...", address)
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:            address,
		Password:        "", // no password set
		DB:              0,  // use default DB
		MaxRetries:      10,
		MinRetryBackoff: 100 * time.Millisecond,
		MaxRetryBackoff: 3 * time.Second, // 15 seconds max
		OnConnect: func(ctx context.Context, conn *redis.Conn) error {
			log.Println("New Redis connection established")
			return nil
		},
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	return RedisClientImpl{client, ctx, key}
}
