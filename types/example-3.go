package types

import "fmt"

func Example3() {
	var x interface{} = 3.14

	switch t := x.(type) {
	case int:
		fmt.Println("x is an int:", t)
	case string:
		fmt.Println("x is a string:", t)
	case float64:
		fmt.Println("x is a float64:", t)
	default:
		fmt.Println("x is of unknown type:", t)
	}
}
