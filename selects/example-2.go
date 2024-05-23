package selects

import (
	"fmt"
	"time"
)

func Example2() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "Deleyed msg"
	}()

	select {
	case <-ch:
		fmt.Println("Received message...")
	default:
		fmt.Println("No msg received...")
	}

	time.Sleep(time.Second * 3)
}
