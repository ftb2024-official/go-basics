package ocp

import (
	"fmt"
	"math"
)

type shape2 interface {
	area() float64
}

type circle2 struct{ radius float64 }
type square2 struct{ length float64 }
type triangle2 struct{ base, height float64 }
type calculator2 struct{}

func (c circle2) area() float64    { return math.Pi * c.radius * c.radius }
func (s square2) area() float64    { return s.length * s.length }
func (tr triangle2) area() float64 { return tr.base * tr.height / 2 }
func (c calculator2) sumArea(shapes ...shape2) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

func Example2() {
	c := circle2{5}
	s := square2{7}
	t := triangle2{3, 7}
	var calc calculator2

	fmt.Printf("total area: %.2f\n", calc.sumArea(c, s, t))
	fmt.Println("height:", t.height)
}
