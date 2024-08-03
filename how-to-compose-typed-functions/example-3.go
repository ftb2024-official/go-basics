package howtocomposetypedfunctions

import "fmt"

func Example3() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filteredNums := filter(nums, isOdd)
	doubleNums := mapp(filteredNums, double)

	fmt.Println("Result:", doubleNums)
}

func isOdd(n int) bool {
	return n%2 != 0
}

func double(n int) int {
	return n * 2
}

func filter(nums []int, predicate func(int) bool) []int {
	var result []int

	for _, num := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}

	return result
}

func mapp(nums []int, fn func(int) int) []int {
	var result []int

	for _, num := range nums {
		result = append(result, fn(num))
	}

	return result
}
