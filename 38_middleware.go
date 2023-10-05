package main

import (
    "fmt"
    "net/http"
)

// Middleware function to log requests
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// Handler function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    // Create a router and add the Logger middleware
    router := http.NewServeMux()
    router.Handle("/", Logger(http.HandlerFunc(HelloWorld)))

    // Start the server
    http.ListenAndServe(":8080", router)
}
