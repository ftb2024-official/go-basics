package closures

import "fmt"

func m() (func(int) int, func(int) int) {
	total := 0

	add := func(x int) int {
		total += x
		return total
	}

	sub := func(x int) int {
		total -= x
		return total
	}

	return add, sub
}

func Example1() {
	add, sub := m()
	fmt.Println(add(5))  // 5
	fmt.Println(sub(3))  // 2
	fmt.Println(sub(5))  // -3
	fmt.Println(add(1))  // -2
	fmt.Println(add(10)) // 8
}
