package mutexplusgoroutines

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock() // захват мьютекса перед изменением переменной balance
	fmt.Printf("Depositing %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock() // освобождение мьютекса после изменения

	wg.Done()
}

func Example1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go deposit(100, &wg)
	go deposit(300, &wg)

	wg.Wait()
	fmt.Println("New balance:", balance)
}
