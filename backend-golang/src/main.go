package main

import (
    "os"
)

const (
    PORT = 8080
)

var redisHost = "localhost"
var greetingLabel = "#greetingLabel#"

func main() {
    if os.Getenv("redisHost") != "" {
        redisHost = os.Getenv("redisHost")
    }
    if os.Getenv("greetingLabel") != "" {
        greetingLabel = os.Getenv("greetingLabel")
    }

    redisClient := NewRedisClient(redisHost)
    defer redisClient.Close()

    httpHandler := NewHttpHandler(redisClient, greetingLabel)
    httpServer := NewHttpServer(httpHandler)
    httpServer.Listen(PORT)
}
