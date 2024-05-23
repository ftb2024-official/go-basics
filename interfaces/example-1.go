package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func printArea(sh Shape) {
	fmt.Println("Area:", sh.Area())
}

func Example1() {
	circle := Circle{13}
	square := Square{21}

	printArea(&circle)
	printArea(&square)
}
