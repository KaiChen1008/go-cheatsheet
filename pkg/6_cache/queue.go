package threadsafe

type Queue[T any] struct {
	ch chan T
}

func New[T any](size int) *Queue[T] {
	return &Queue[T]{
		ch: make(chan T, size),
	}
}

func (q *Queue[T]) Eunqueue(v T) {
	q.ch <- v
}

func (q *Queue[T]) Dequeue() T {
	return <-q.ch
}
