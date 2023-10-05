package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(i)
	}

	// Send jobs to workers
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	// Collect results from workers
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
