package waitgroups

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done() // вызов 'Done' в 'defer' гарантирует, что счетчик уменьшится при выходе из функции

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second) // имитируем выполнение какой-либо работы
	fmt.Printf("Worker %d finished\n", id)
}

func Example1() {
	var wg sync.WaitGroup
	totalWorkers := 10

	wg.Add(totalWorkers) // установка счетчика на кол-во работников
	for i := 1; i <= totalWorkers; i++ {
		go worker(&wg, i) // запуск каждого работника в отдельной горутине
	}

	wg.Wait() // ждем завершения всех горутин
	fmt.Println("All workers finished their job")
}
