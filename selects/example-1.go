/*
Используем 'select' для чтения данных из двух каналов. Мы будем отправлять сообщения в каналы из разных горутин и использовать 'select', чтобы
обработать первое доступное сообщение.
*/

package selects

import (
	"fmt"
	"time"
)

func Example1() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// горутина для отправки сообщения в первый канал через 1 секунду
	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- "Msg from ch1"
	}()

	// горутина для отправки сообщения во второй канал через 2 секунды
	go func() {
		time.Sleep(time.Second * 5)
		ch2 <- "Msg from ch2"
	}()

	// используем 'select' для ожидания сообщений из двух каналов
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received msg from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received msg from ch2:", msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("Timeout occured, no data recieved within 3 seconds...")
		}
	}
}
