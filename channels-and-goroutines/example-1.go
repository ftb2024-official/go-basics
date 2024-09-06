package channelsandgoroutines

import "fmt"

// функция для вычисления факториала числа 'n'
func factorial(n int, resultChan chan<- int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	resultChan <- result // отправляем результат в канал
}

// функция для вывода результатов из канала
func printResults(resultCh <-chan int, doneChan chan<- bool, count int) {
	for i := 0; i < count; i++ {
		result := <-resultCh // получаем результат из канала
		fmt.Println(result)
	}

	doneChan <- true // сообщаем, что все результаты были обработаны
}

func Example1() {
	nums := []int{5, 7, 10, 12, 15}
	resultChan := make(chan int, len(nums))
	doneChan := make(chan bool)

	for _, num := range nums {
		go factorial(num, resultChan) // запускаем горутину для вычисления факториала
	}

	go printResults(resultChan, doneChan, len(nums)) // запускаем горутину для выводы результата

	// ждем пока обработаются все результаты
	fmt.Println("All results are processed:", <-doneChan)
}
