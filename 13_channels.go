package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a channel for integers
    ch := make(chan int)

    go func() {
        // Send the value 42 into the channel
        ch <- 42
    }()

    // Receive and print the value from the channel
    value := <-ch
    fmt.Println(value)
}
// Expected Output: 42
