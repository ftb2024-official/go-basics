package slices

import (
	"fmt"
	"sync"
)

func incSlice(slice []int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for i := range slice {
		mu.Lock()
		slice[i]++
		mu.Unlock()
	}
}

func Example2() {
	nums := []int{1, 2, 3, 4, 5}
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	wg.Add(2)

	go incSlice(nums, &wg, &mu)
	go incSlice(nums, &wg, &mu)

	wg.Wait()
	fmt.Println(nums)
}
