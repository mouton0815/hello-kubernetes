package main

import (
    "github.com/gomodule/redigo/redis"
    "log"
)

type RedisClient interface {
    Incr(key string) int
    Close()
}

type RedisClientImpl struct {
    conn redis.Conn
}

func (r RedisClientImpl) Incr(key string) int {
    res, err := redis.Int(r.conn.Do("INCR", key))
    if err != nil {
        log.Fatal(err)
    }
    return res
}

func (r RedisClientImpl) Close() {
    _ = r.conn.Close()
}

// Factory function
func NewRedisClient(host string) RedisClient {
    address := host + ":6379"
    log.Printf("Connecting to Redis at %s ...", address)
    conn, err := redis.Dial("tcp", address)
    if err != nil {
        log.Fatal(err)
    }
    return RedisClientImpl{ conn }
}
