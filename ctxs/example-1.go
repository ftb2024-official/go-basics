package ctxs

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err   error
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 100)
	return 777, nil
}

func fetchUserData(ctx context.Context) (int, error) {
	ctxVal := ctx.Value("foo")
	fmt.Println("context value:", ctxVal)

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200) // will wait only fot 200ms
	defer cancel()

	respCh := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respCh <- Response{value: val, err: err}
	}()

	for {
		select {
		case <-ctx.Done(): // if nothing happened in 200ms
			return 0, fmt.Errorf("fetching data from third party took too long...")
		case resp := <-respCh: // if we got response from third party api
			return resp.value, resp.err
		}
	}
}

func Example1() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "foo", "bar")
	val, err := fetchUserData(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))
}
