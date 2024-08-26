package closures

import "fmt"

func createMultiplier(m int) func(n int) int {
	mult := func(n int) int {
		return m * n
	}
	return mult
}

func Example2() {
	multBy2 := createMultiplier(2)
	multBy5 := createMultiplier(5)

	fmt.Println(multBy2(7))  // 14
	fmt.Println(multBy5(10)) // 50
}
