package main

import (
    "fmt"
    "time"
)

func printHello() {
    // Print "Hello" in a Goroutine
    fmt.Println("Hello")
}

func main() {
    // Launch a Goroutine to print "Hello" concurrently
    go printHello()
    time.Sleep(1 * time.Second)
}
// Expected Output: Hello
