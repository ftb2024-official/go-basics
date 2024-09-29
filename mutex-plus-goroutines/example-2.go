package mutexplusgoroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func Example2() {
	var (
		readOps, writeOps uint64
		state             = make(map[int]int)
		mutex             = &sync.Mutex{}
	)

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)

				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)

				mutex.Lock()
				state[key] = value
				mutex.Unlock()

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	fmt.Println("readOps:", readOpsFinal)
	fmt.Println("writeOps:", writeOpsFinal)
}
