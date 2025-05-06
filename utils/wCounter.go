package utils

import "fmt"

func WCounter(n int) (int, error) {
	if n <= 0 {
		return -1, fmt.Errorf("n must be greater than 0")
	}

	if n == 1 {
		return 4, nil
	}

	total := n*3 + 2
	for i := 2; i < n; i++ {
		total += i * 4
	}

	return total, nil
}
