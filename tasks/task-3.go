package tasks

import (
	"fmt"
	"sync"
)

func Task3() {
	var wg sync.WaitGroup
	m := make(chan string, 3)
	chNum := 11

	for i := 0; i < chNum; i++ {
		wg.Add(1)
		go func(v int) {
			m <- fmt.Sprintf("Goroutine %d", v)
		}(i)
	}

	for i := 0; i < chNum; i++ {
		go receiveFromCh(m, &wg)
	}

	wg.Wait()
}

func receiveFromCh(ch chan string, wg *sync.WaitGroup) {
	fmt.Println(<-ch)
}
