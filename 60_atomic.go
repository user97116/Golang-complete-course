package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32
	var wg sync.WaitGroup

	numGoroutines := 100

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Counter Value:", atomic.LoadInt32(&counter))
}
