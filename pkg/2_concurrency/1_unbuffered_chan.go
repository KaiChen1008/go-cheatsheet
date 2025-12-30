package concurrency

func main() {
	ch := make(chan string) // unbuffered

	go push("Moe", ch)
	go push("Larry", ch)
	go push("Claire", ch)

	// goroutines are concurrent, the order isn't guaranteed.
	println(<-ch)
	println(<-ch)
	println(<-ch)
}

func push(name string, ch chan<- string) {
	ch <- "Hey" + name
}
