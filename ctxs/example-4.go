/*
Иногда нужно явно отменить контекст, например, при получении сигнала от пользователя или другой части программы.
*/

package ctxs

import (
	"context"
	"fmt"
	"time"
)

func Example4() {
	// родительский контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем горутину передавая контекст
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine: Context cancelled...")
				return
			default:
				fmt.Println("Goroutine: Working...")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	// ждем 2 секунды, перед отменой контекста
	time.Sleep(time.Second * 5)
	cancel() // отмена контекста

	// ждем секунду, чтобы горутина завершилась
	time.Sleep(time.Second)
	fmt.Println("Main: Done...")
}
