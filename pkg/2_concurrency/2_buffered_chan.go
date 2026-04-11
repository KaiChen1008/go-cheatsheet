package concurrency

func BufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // fatal error: all goroutines are asleep -> deadlock!
}

/*
The loop for i := range c receives values from the channel repeatedly until it is closed.

Note 1: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
Note 2: Channels aren't like files; you don't usually need to close them.

	Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
*/
func CloseChannel() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for i := range ch {
		println(i)
	}
	// method 2:
	// v, ok := <- ch
}

/*
what will happend if we get data from a closed channel -> repeatedly receive (default val, false)
output:

	0 false
	0 false
	0 false
	...
*/
func CloseChannel2() {
	q := make(chan int, 10)
	close(q)
	for range 10 {
		v, ok := <-q
		println(v, ok)
	}
}
