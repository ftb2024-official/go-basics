package solid

import (
	"fmt"
	"math"
)

type shape4 interface {
	area() float64
}

type circle4 struct{ radius float64 }
type square4 struct{ length float64 }
type triangle4 struct{ base, height float64 }
type calculator4 struct{}

func (c circle4) area() float64    { return math.Pi * c.radius * c.radius }
func (s square4) area() float64    { return s.length * s.length }
func (tr triangle4) area() float64 { return tr.base * tr.height / 2 }
func (c calculator4) sumArea(shapes ...shape4) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

func Example4() {
	c := circle4{5}
	s := square4{7}
	t := triangle4{3, 7}
	var calc calculator4

	fmt.Printf("total area: %.2f\n", calc.sumArea(c, s, t))
	fmt.Println("height:", t.height)
}
