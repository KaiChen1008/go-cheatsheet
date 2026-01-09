package racecondition

import (
	"fmt"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	sem := New(3)

	for i := range 10 {
		go func(id int) {
			sem.Acquire()
			defer sem.Release()

			fmt.Printf("processing task %d\n", id)
			time.Sleep(1 * time.Second) // Simulate work
		}(i)
	}

	time.Sleep(4 * time.Second) // Ensure all goroutines complete
}
