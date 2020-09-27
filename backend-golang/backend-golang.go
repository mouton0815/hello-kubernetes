//
// Example taken from https://hackernoon.com/how-to-create-a-web-server-in-go-a064277287c9
// 

package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"
)

const (
    PORT = "8080"
)

var greetingLabel string = "#greetingLabel#"

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
    message = "[GOLANG] " + greetingLabel + " " + strings.TrimPrefix(message, "/")
    _, err := w.Write([]byte(message))
    if err != nil {
        fmt.Printf("Error when sending response %v", err)
    }
}

func main() {
    if os.Getenv("greetingLabel") != "" {
        greetingLabel = os.Getenv("greetingLabel")
    }
    mux := http.NewServeMux()
    mux.HandleFunc("/", sayHello)
    fmt.Printf("Server listens on port %s\n", PORT)
    if err := http.ListenAndServe(":" + PORT, RequestLogger(mux)); err != nil {
        panic(err)
    }
}
