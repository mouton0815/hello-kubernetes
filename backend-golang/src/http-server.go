package main

import (
    "log"
    "net/http"
    "strconv"
)

type HttpServer interface {
    Listen(port int)
}

type HttpServerImpl struct {
    server http.Handler
}

func (h HttpServerImpl) Listen(port int) {
    log.Printf("Server listens on port %d\n", port)
    if err := http.ListenAndServe(":" + strconv.Itoa(port), HttpLogger(h.server)); err != nil {
        panic(err)
    }
}

func NewHttpServer(handler HttpHandler) HttpServer {
    server := http.NewServeMux()
    server.HandleFunc("/", handler.Hello)
    return HttpServerImpl{ server }
}
