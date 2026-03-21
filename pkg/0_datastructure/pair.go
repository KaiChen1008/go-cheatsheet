package datastructure

/*
In Go, structs are comparable if all their fields are comparable types.
*/
type pair struct {
	x, y int
}

func main() {
	a := pair{1, 2}
	b := pair{1, 2}
	println(a == b) // true
}
