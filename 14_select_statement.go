package main

import (
    "fmt"
    "time"
)

func main() {
    // Create two channels
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- 1
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- 2
    }()

    // Use the select statement to receive from the first available channel
    select {
    case value := <-ch1:
        fmt.Println("Received from ch1:", value)
    case value := <-ch2:
        fmt.Println("Received from ch2:", value)
    }
}
// Expected Output: Received from ch2: 2
