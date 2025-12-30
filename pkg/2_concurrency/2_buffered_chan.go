package concurrency

func BufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // fatal error: all goroutines are asleep -> deadlock!
}

func CloseAChannel() {
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
The loop for i := range c receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
*/
