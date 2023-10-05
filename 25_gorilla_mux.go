package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Create a simple RESTful API using Gorilla Mux
    router := mux.NewRouter()

    // Define a handler for the /hello/{name} route
    router.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        name := vars["name"]
        fmt.Fprintf(w, "Hello, %s!", name)
    })

    // Start the server on port 8080
    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}
// Expected Output: Access "http://localhost:8080/hello/John" in a web browser to see "Hello, John!"
