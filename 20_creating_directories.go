package main

import (
    "fmt"
    "os"
)

func main() {
    // Create a new directory
    err := os.Mkdir("mydir", 0755)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Directory created successfully.")
}
// Expected Output: Directory created successfully.
