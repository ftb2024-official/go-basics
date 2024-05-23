package ctxs

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func requestRide(ctx context.Context, serviceName string, resChan chan string) {
	time.Sleep(time.Second * 3)

	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped the search in %q (%v)", serviceName, ctx.Err())
			return
		default:
			if rand.Float64() > 0.75 {
				resChan <- serviceName
				return
			}
		}

		continue
	}
}

func CallRequestRide() {
	var (
		resChan     = make(chan string)
		services    = []string{"Super", "VillageMobil", "Set Taxi", "Index Go"}
		ctx, cancel = context.WithCancel(context.Background())
		winner      string
		wg          sync.WaitGroup
	)

	defer cancel()

	for i := range services {
		svc := services[i]

		wg.Add(1)
		go func() {
			requestRide(ctx, svc, resChan)
			wg.Done()
		}()
	}

	go func() {
		winner = <-resChan
		cancel()
	}()

	wg.Wait()
	log.Println("found car in", winner)
}
