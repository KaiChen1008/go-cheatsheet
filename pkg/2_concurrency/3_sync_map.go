package concurrency

import (
	"fmt"
	"sync"
)

// ref: https://go-cookbook.com/snippets/concurrency/concurrency-and-data-structures
func SyncMap() {
	sm := sync.Map{}
	wg := sync.WaitGroup{}
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("task%d", i)

			// sm.Store(key, i)
			if _, loaded := sm.LoadOrStore(key, i); loaded { // use LoadOrStore to handle race conditions.
				println("key exists")
			}
		}(i)
	}

	wg.Wait()

	sm.Range(func(key, value any) bool {
		fmt.Printf("%s: %d\n", key, value)
		return true
	})
}
