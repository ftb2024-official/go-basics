package main

import "fmt"

func f1(n int) float64 {
	result := 10_000.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result
}

func f2(n int) float64 {
	result := 0.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result + 10_000.
}

func main() {
	v1 := f1(1000)
	v2 := f2(1000)

	fmt.Println(v1)
	fmt.Println(v2)
}
