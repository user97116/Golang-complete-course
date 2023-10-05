package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Create a simple HTTP server that responds with "Hello, World!"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    })

    // Start the server on port 8080
    http.ListenAndServe(":8080", nil)
}
// Expected Output: Access "http://localhost:8080" in a web browser to see "Hello, World!"
