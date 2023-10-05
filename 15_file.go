package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    // Define content to write to a file
    content := []byte("Hello, File!")

    // Write the content to a file named "example.txt"
    err := ioutil.WriteFile("example.txt", content, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("File created and content written successfully.")
}
// Expected Output: File created and content written successfully.
