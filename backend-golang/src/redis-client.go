package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type RedisClient interface {
	Incr() int
	Close()
}

type RedisClientImpl struct {
	conn redis.Conn
	key  string
}

func (r RedisClientImpl) Incr() int {
	res, err := redis.Int(r.conn.Do("INCR", r.key))
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (r RedisClientImpl) Close() {
	_ = r.conn.Close()
}

// NewRedisClient Factory function
func NewRedisClient(host string, key string) RedisClient {
	address := host + ":6379"
	log.Printf("Connecting to Redis at %s ...", address)
	// Redis may not be up and running yet, so try at most three times to connect
	attempts := 1
	for {
		client, err := createRedisClient(address, key)
		if err == nil {
			log.Println("Connection to Redis established - attempt", attempts)
			return client
		}
		log.Println(err, "- attempt", attempts)
		if attempts >= 3 {
			panic(err)
		}
		var duration time.Duration = 3 * 1000 * 1000 * 1000 // 3 seconds
		time.Sleep(duration)
		attempts += 1
	}
}

func createRedisClient(address string, key string) (RedisClient, error) {
	var timeout time.Duration = 3 * 1000 * 1000 * 1000 // 3 seconds
	conn, err := redis.Dial("tcp", address, redis.DialConnectTimeout(timeout))
	if err == nil {
		return RedisClientImpl{conn, key}, nil
	}
	return nil, err
}
