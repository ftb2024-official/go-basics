package defers

import "fmt"

func Example2() {
	x := 10

	defer func(val int) {
		fmt.Println("value of x(copy):", val) // 10
	}(x)

	defer func() {
		fmt.Println("value of x with closure", x) // 20
	}()

	defer fmt.Println("value of x with defer:", x) // 10

	x = 20
	fmt.Println("value of x:", x) // 20
}
