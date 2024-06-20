package slices

import (
	"fmt"
	"sync"
)

func printSlice(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, v := range slice {
		fmt.Println(v)
	}
}

func Example1() {
	nums := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	wg.Add(2)

	go printSlice(nums, &wg)
	go printSlice(nums, &wg)

	wg.Wait()
}
