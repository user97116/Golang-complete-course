package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup

    // Create and start multiple workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    // Wait for all workers to finish
    wg.Wait()

    fmt.Println("All workers have finished.")
}
// Expected Output:
// Worker 1 started
// Worker 2 started
// Worker 3 started
// Worker 1 finished
// Worker 2 finished
// Worker 3 finished
// All workers have finished.
