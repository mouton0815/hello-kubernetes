package main

import (
	"os"
)

const (
	PORT = 8080
)

func main() {
	redisHost := getEnv("redisHost", "localhost")
	greetingLabel := getEnv("greetingLabel", "#greetingLabel#")

	redisClient := NewRedisClient(redisHost, "backend-golang-counter")
	defer redisClient.Close()

	httpHandler := NewHttpHandler(redisClient, greetingLabel)
	httpServer := NewHttpServer(httpHandler)
	httpServer.Listen(PORT)
}

func getEnv(name string, fallback string) string {
	value, ok := os.LookupEnv(name)
	if ok {
		return value
	}
	return fallback
}
