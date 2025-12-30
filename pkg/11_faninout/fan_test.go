package faninout

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFanIn2(t *testing.T) {
	goroutines := runtime.NumGoroutine()

	producer := func(ch chan int) {
		defer close(ch)
		for i := range 10 {
			ch <- i
		}
	}

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go producer(ch1)
	go producer(ch2)

	for v := range FanIn(ch1, ch2) {
		println(v)
	}
	require.Equal(t, goroutines, runtime.NumGoroutine())
}
