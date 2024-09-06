package channelsandgoroutines

import "fmt"

func squares(c chan int) {
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func Example5() {
	fmt.Println("main() started...")
	chInt := make(chan int, 3)

	go squares(chInt)

	chInt <- 1
	chInt <- 2
	chInt <- 3
	chInt <- 4 // goroutine 'main' will be blocked here until another goroutine reads data from the channel

	fmt.Println("main() stopped...")
}
