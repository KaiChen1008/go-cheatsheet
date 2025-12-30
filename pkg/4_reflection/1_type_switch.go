package reflection

import "fmt"

func Reflection() {
	checkType(21)
	checkType("hello")
	checkType(true)
}

func checkType(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is int\n", i)
	case string:
		fmt.Printf("%v is string, %v\n", i, v)
	default:
		fmt.Printf("unknown type\n")
	}
}
