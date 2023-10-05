package main

import (
    "fmt"
    "os"
)

func main() {
    // Open a file for appending
    file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Append content to the file
    _, err = file.WriteString("\nAppended Text")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("File appended successfully.")
}
// Expected Output: File appended successfully.
