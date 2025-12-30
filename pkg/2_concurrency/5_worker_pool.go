package concurrency

import (
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers = 3
	numJobs    = 10
)

func worker(wg *sync.WaitGroup, id int, jobs <-chan int, results chan<- int) {
	defer wg.Done()

	// use range chan to receive jobs
	for job := range jobs {
		results <- job * 2
		fmt.Printf("job %v is processed by worker %v\n", job, id)
		time.Sleep(time.Millisecond)
	}
}

func Run() {
	jobs := make(chan int, 5)
	results := make(chan int, numJobs)

	wg := sync.WaitGroup{}
	// start workers
	for i := range numWorkers {
		i := i
		wg.Add(1)
		go worker(&wg, i, jobs, results)
	}

	// send jobs
	for i := range numJobs {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("Result:", r)
	}
}
