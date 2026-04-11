package racecondition

import (
	"fmt"
	"sync"
)

// use go run -race main.go to dectect race conditions

func main() {
	var value int
	var wg sync.WaitGroup

	for range 10 {
		wg.Go(func() {
			value++
		})
	}

	wg.Wait()
	fmt.Println("Value:", value)
}
