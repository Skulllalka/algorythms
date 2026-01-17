package solutions

import (
	"fmt"
	"sync"
)

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- f(j)
	}
}

const (
	numJobs    = 5
	numWorkers = 3
)

func WorkerPoolMain() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numWorkers)
	wg := sync.WaitGroup{}

	multiplier := func(x int) int {
		return x * x
	}

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(multiplier, jobs, results)
		}()
	}

	done := make(chan struct{})
	go func() {
		for r := range results {
			fmt.Println(r)
		}
		close(done)
	}()

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	<-done
}
