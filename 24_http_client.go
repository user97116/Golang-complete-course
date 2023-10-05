package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    // Make an HTTP GET request and print the response
    response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("HTTP Response:")
    fmt.Println(string(body))
}
// Expected Output: HTTP Response: {"userId":1,"id":1,"title":"...",
