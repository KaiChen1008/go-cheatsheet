package datastructure

type pair struct {
	x, y int
}

func main() {
	a := pair{1, 2}
	b := pair{1, 2}
	println(a == b)
}
