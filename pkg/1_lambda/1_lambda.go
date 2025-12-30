package lambda

func Simple() {
	x := 10
	myFunc := func() bool {
		return x > 1000
	}
	myFunc()
}

func WithIf() {
	x := 100
	if ok := func() bool {
		return x > 1000
	}(); !ok {
		println("failed")
	}
	print("success")
}
