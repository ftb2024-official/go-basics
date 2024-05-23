package ctxs

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) (string, error) {
	select {
	case <-time.After(time.Second * 1):
		return "Some result", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func Example5() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	if result, err := longRunningOperation(ctx); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
