package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // Read and print the names of files in a directory
    files, err := ioutil.ReadDir(".")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    for _, file := range files {
        fmt.Println("File:", file.Name())
    }
}
// Expected Output: File: example.txt
