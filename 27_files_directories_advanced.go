package main

import (
    "fmt"
    "os"
)

func main() {
    // Create a directory and a file within it
    err := os.MkdirAll("mydir/subdir", 0755)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    file, err := os.Create("mydir/file.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    fmt.Println("Directory and file created successfully.")
}
// Expected Output: Directory and file created successfully.
