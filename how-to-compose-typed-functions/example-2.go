package howtocomposetypedfunctions

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func compose(f, g func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		return f(g(x, y), y)
	}
}

func Example2() {
	addThenMultiply := compose(multiply, add)
	result := addThenMultiply(2, 3) // (2 + 3) * 3 = 15
	fmt.Println("Result:", result)
}
