package main

import (
    "fmt"
    "os"
)

func main() {
    // Retrieve and print the value of an environment variable
    homeDir := os.Getenv("HOME")
    fmt.Println("Home Directory:", homeDir)
}
// Expected Output: Home Directory: /home/username
