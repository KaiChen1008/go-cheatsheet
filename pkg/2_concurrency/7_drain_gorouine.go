package concurrency

import "time"

func DrainGoroutine() {
	ch := make(chan struct{}, 10)
	ch <- struct{}{}
	ch <- struct{}{}
	close(ch)

	timer := time.NewTimer(time.Millisecond)
	for {
		select {
		case <-timer.C:
			return
		case _, ok := <-ch:
			/*
				當 channel 被 close 後，所有的 receiver 會“一直”收到 0, false
				此時，這個 case 會一直執行下去 -> drain goroutine
				所以要判斷 !ok -> return
			*/
			if !ok {
				return
			}
			println("hi")
		}
	}
}
