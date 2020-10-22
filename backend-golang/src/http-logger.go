package main

import (
    "fmt"
    "net/http"
    "time"
)

func HttpLogger(targetMux http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        targetMux.ServeHTTP(w, r)
        requesterIP := r.RemoteAddr
        fmt.Printf(
            "%s %s %s %v\n",
            r.Method,
            r.RequestURI,
            requesterIP,
            time.Since(start),
        )
    })
}

