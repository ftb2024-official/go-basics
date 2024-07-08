package types

import "fmt"

func Example2() {
	var x interface{} = "hello, world"

	if str, ok := x.(string); ok {
		fmt.Println("x is a string:", str)
	} else {
		fmt.Println("x isn't a string:", x)
	}
}
