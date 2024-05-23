package ctxs

import (
	"context"
	"fmt"
	"time"
)

func Example3() {
	// создаем родительский контекст с тайм-аутом 3 секунды
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel() // отмена контекста по завершении 'main'

	// запускаем горутину, передавая ей контекст
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine: Context cancelled or timed out")
				return
			default:
				fmt.Println("Goroutine: Working...")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	// ждем 5 секунд перед завершением 'main'
	time.Sleep(time.Second * 5)
	fmt.Println("Main: Done")
}
