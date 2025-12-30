package racecondition

type Semaphore struct {
	counter chan struct{}
}

func New(size int) *Semaphore {
	return &Semaphore{
		counter: make(chan struct{}, size),
	}
}

func (s *Semaphore) Acquire() {
	// block when the chan is full
	s.counter <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.counter
}

// do not close the channel used in a semaphore,
// as this is not necessary and could lead to unintended panics if more goroutines continue to operate on the semaphore.
