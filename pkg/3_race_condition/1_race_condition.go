package racecondition

import (
	"sync"
	"sync/atomic"
)

var counter int64

func Inc() {
	counter++ // race
}

// correct way 1
type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

// correct way 2
func IncAotmic() {
	atomic.AddInt64(&counter, 1)
}
