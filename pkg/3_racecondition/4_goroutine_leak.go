package racecondition

import (
	"context"
	"time"
)

func Worker(ctx context.Context, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return // leak
		case job := <-jobs:
			go process(job) // <- goroutine does not finished
		}
	}
}

func process(_ int) {
	time.Sleep(5 * time.Second)
}
