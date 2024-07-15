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

type Rectangle struct {
	Length, Width float64
}

type Square struct {
	Side float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func printArea(sh Shape) {
	fmt.Printf("%.2f\n", sh.Area())
}

func totalArea(shapes []Shape) float64 {
	var total float64

	for _, sh := range shapes {
		total += sh.Area()
	}

	return total
}

func Example1() {
	c := Circle{10}
	r := Rectangle{20, 30}
	s := Square{40}

	shapes := []Shape{c, r, s}

	printArea(c)
	printArea(r)
	printArea(s)
	fmt.Printf("Total area: %.2f\n", totalArea(shapes))
}
