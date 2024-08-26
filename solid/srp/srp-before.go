// Single Responsibility Principle
package srp

import (
	"fmt"
	"math"
)

type circle struct{ radius float64 }

func (c circle) area() { fmt.Printf("circle area: %.2f\n", math.Pi*c.radius*c.radius) }

type square struct{ length float64 }

func (s square) area() { fmt.Printf("square area: %.2f\n", s.length*s.length) }

func Example1() {
	c := circle{5}
	s := square{7}

	c.area()
	s.area()
}
