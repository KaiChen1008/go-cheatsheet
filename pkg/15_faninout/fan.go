package faninout

import (
	"sync"
)

// fan in: aggregate data from multiple channels to one channel
func FanIn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// fan out: deliver jobs to multiple workers
func FanOut() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 100 {
			ch <- i
		}
	}()

	worker := func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := range ch {
			println(i)
		}
	}

	for range 10 {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}
