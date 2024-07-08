/*
Использование каналов для управления горутинами
*/

package channelsplusgoroutines

import (
	"fmt"
	"sync"
	"time"
)

func workerWithChan(stopCh <-chan struct{}, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			fmt.Printf("Worker %d stopping...\n", id)
			return
		default:
			// выполнение полехной работы
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func Example3() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})

	// запуск горутин
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go workerWithChan(stopCh, &wg, i)
	}

	// остановка горутин после 3 секунд
	time.Sleep(time.Second * 3)
	close(stopCh) // отправка сигнала всем горутинам остановиться
	wg.Wait()     // ожидание завершение всех горутин
}
