package howtousethectx

import (
	"context"
	"errors"
	"log"
	"time"
)

func Example1() {
	start := time.Now()
	ctx := context.WithoutCancel(context.Background())
	result, err := funcToCallSlowFunc(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("result:", result)
	log.Println("took us:", time.Since(start))
}

func verySlowFunc() (int, error) {
	time.Sleep(time.Millisecond * 100)
	return 777, nil
}

func funcToCallSlowFunc(ctx context.Context) (int, error) {
	ctx200, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	val, err := verySlowFunc()
	if err != nil {
		return 0, err
	}

	for {
		select {
		case <-ctx200.Done():
			return 0, errors.New("can't wait for response from slow server")
		default:
			return val, nil
		}
	}
}
