package concurrency

import "sync"

var itemList = []string{
	"Claire",
	"Amy",
	"Ken",
}

func WaitGroup() {
	wg := sync.WaitGroup{}
	for _, item := range itemList {
		wg.Add(1)
		do(&wg, item)
	}

	wg.Wait()
}

func do(wg *sync.WaitGroup, item string) {
	defer wg.Done()
	println(item)
}
