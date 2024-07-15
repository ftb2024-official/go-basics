package functions

import "fmt"

func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func Example1() {
	a, b := 5, 10

	sum := applyOperation(a, b, add)
	fmt.Println("Sum:", sum)

	product := applyOperation(a, b, multiply)
	fmt.Println("Product:", product)
}
