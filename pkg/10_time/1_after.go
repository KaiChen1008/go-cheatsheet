package time

import "time"

// pitfal
func After(ch <-chan int) {
	// create a timer outside a for loop
	// it will repeatly count down a second
	timer := time.NewTimer(time.Second)
	defer timer.Stop()

	for {
		select {
		case <-ch:
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(time.Second)
		case <-timer.C:
			return

		// pitfall 1: time.After create an timer every time
		// pitfall 2: "break" only break the "select" instead of the for loop
		case <-time.After(time.Second):
			break
		}
	}
}

func After2(ch <-chan int) {
	// case 2, only timeout one time
	select {
	case <-ch:
	case <-time.After(time.Second):
		return
	}
}

// Overusing time.After inside a loop which can lead to memory leaks, as it creates a new timer on each iteration.
