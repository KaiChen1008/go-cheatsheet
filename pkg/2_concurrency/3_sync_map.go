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

		// pitfall fixed: https://medium.com/@krisguttenbergovitz/go-1-22s-loop-variable-fix-solving-a-decade-old-gotcha-and-modern-concurrency-pitfalls-69aa4eb0b8a1
		// no need i := i or fun(i int) {...}(i)

		go func() {
			defer wg.Done()

			// sm.Store(i, i)
			if _, loaded := sm.LoadOrStore(i, i); loaded { // use LoadOrStore to handle race conditions.
				println("key exists")
			}
		}()
	}

	wg.Wait()

	sm.Range(func(key, value any) bool {
		fmt.Printf("%s: %d\n", key, value)
		return true
	})
}
