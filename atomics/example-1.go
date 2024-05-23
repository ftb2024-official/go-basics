package atomics

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Example1() {
	var counter int64 // счетчик, который будем увеличивать атомарно
	var wg sync.WaitGroup

	numOfGoroutines := 100   // кол-во горутин
	incsPerGoroutine := 1000 // кол-во инкрементов на одну горутину

	for i := 0; i < numOfGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incsPerGoroutine; j++ {
				atomic.AddInt64(&counter, 1) // атомарный инкремент
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

func Example2() {
	var counter int64
	var wg sync.WaitGroup
	var mutex sync.Mutex

	numOfGoroutines := 200
	incsPerGoroutine := 1000

	for i := 0; i < numOfGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incsPerGoroutine; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
