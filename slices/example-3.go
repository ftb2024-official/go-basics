package slices

import (
	"fmt"
	"sync"
)

func incrementSlice(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range slice {
		slice[i]++
	}
}

func Example3() {
	slice := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	wg.Add(2)

	// передаем копии слайса в горутины
	go incrementSlice(append([]int(nil), slice...), &wg)
	go incrementSlice(append([]int(nil), slice...), &wg)

	wg.Wait()
	fmt.Println(slice)
}
