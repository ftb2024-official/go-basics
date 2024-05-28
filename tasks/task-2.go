package tasks

import "fmt"

func Task2() {
	var numbers []*int
	for _, value := range []int{10, 20, 30, 40} {
		numbers = append(numbers, &value)
	}

	for _, num := range numbers {
		fmt.Printf("%v ", *num)
	}
}
