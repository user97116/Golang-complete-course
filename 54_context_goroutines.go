package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Exiting\n", id)
			return
		default:
			fmt.Printf("Worker %d: Working\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Create and start multiple workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// Simulate a condition that triggers cancellation
	time.Sleep(3 * time.Second)
	cancel()

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers have exited.")
}
