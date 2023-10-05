package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // Read and print the contents of a file
    content, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("File Content:\n%s", content)
}
// Expected Output: File Content: Hello, File!
