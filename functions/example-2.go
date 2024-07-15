package functions

import "fmt"

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func Example2() {
	double := createMultiplier(2)
	fmt.Println("Double 5:", double(5))

	triple := createMultiplier(3)
	fmt.Println("Triple 7:", triple(7))
}
