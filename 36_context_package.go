package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
    defer cancel()

    select {
    case <-time.After(5 * time.Second):
        fmt.Fprintln(w, "Hello, World!")
    case <-ctx.Done():
        fmt.Fprintln(w, "Request canceled or timed out.")
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
// Access "http://localhost:8080" in a web browser to observe the timeout response.
