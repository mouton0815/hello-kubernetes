package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

type HttpHandler interface {
	Hello(w http.ResponseWriter, r *http.Request)
}

type HttpHandlerImpl struct {
	redisClient   RedisClient
	greetingLabel string
}

func (h HttpHandlerImpl) Hello(w http.ResponseWriter, r *http.Request) {
	count := strconv.Itoa(h.redisClient.Incr())
	name := strings.TrimPrefix(r.URL.Path, "/")
	message := "[GOLANG] " + h.greetingLabel + " " + name + " (call #" + count + ")"
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Printf("Error when sending response %v", err)
	}
}

func NewHttpHandler(redisClient RedisClient, greetingLabel string) HttpHandler {
	return HttpHandlerImpl{redisClient, greetingLabel}
}
