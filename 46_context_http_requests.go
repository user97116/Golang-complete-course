package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    select {
    case <-time.After(2 * time.Second):
        fmt.Fprintln(w, "Request processed successfully.")
    case <-ctx.Done():
        err := ctx.Err()
        fmt.Fprintln(w, "Request canceled:", err)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
// Access "http://localhost:8080" in a web browser to observe the timeout response.
