/*
Использование пакета context
*/

package channelsandgoroutines

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerWithCtx(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping...\n", id)
			return
		default:
			// выполнение полехной работы
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func Example4() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	// запуск горутин
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go workerWithCtx(ctx, &wg, i)
	}

	wg.Wait()
}
