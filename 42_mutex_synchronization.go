package main

import (
    "fmt"
    "sync"
    "time"
)

var counter int
var mu sync.Mutex

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

func main() {
    // Create multiple Goroutines to increment the counter
    for i := 0; i < 10; i++ {
        go increment()
    }

    // Wait for the Goroutines to finish
    time.Sleep(time.Second)

    fmt.Printf("Counter value: %d\n", counter)
}
// Expected Output: Counter value: 10
