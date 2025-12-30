package time

import (
	"errors"
	"time"
)

// pitfal
func After(ch <-chan int) (int, error) {
	for {
		select {
		case v := <-ch:
			return v, nil
		// pitfall 1: overusing time.After inside a loop which can lead to memory leaks, as it creates a new timer on each iteration
		// pitfall 2: "break" only break the "select" instead of the for loop
		case <-time.After(time.Second):
			break
		}
	}
}

// correct way 1: use time.After without for loop
func After2(ch <-chan int) (int, error) {
	// case 2, only timeout one time
	select {
	case v := <-ch:
		return v, nil
	case <-time.After(time.Second):
		return 0, errors.New("timeout")
	}
}

// correct way 2: use timer and label to break
// this func acts weird, just an example to show how to use labels and timers
func After3(ch <-chan int) (int, error) {
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
Loop:
	for {
		select {
		case v := <-ch:
			return v, nil
		case <-timer.C:
			break Loop
		default:
			println("default")
		}
	}
	return 0, errors.New("timeout")
}
