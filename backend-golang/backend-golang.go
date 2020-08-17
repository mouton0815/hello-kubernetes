//
// Example taken from https://hackernoon.com/how-to-create-a-web-server-in-go-a064277287c9
// 

package main

import (
    "fmt"
    "net/http"
    "strings"
    "time"
)

const (
    PORT = "8080"
)

func RequestLogger(targetMux http.Handler) http.Handler {
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

func sayHello(w http.ResponseWriter, r *http.Request) {
    message := r.URL.Path
    message = "[GOLANG] Hello " + strings.TrimPrefix(message, "/")
    _, err := w.Write([]byte(message))
    if err != nil {
        fmt.Printf("Error when sending response %v", err)
    }
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", sayHello)
    fmt.Printf("Server listens on port %s\n", PORT)
    if err := http.ListenAndServe(":" + PORT, RequestLogger(mux)); err != nil {
        panic(err)
    }
}
