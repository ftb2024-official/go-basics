package solid

import (
	"fmt"
	"math"
)

type circle3 struct{ radius float64 }
type square3 struct{ length float64 }
type calculator struct{}

func (c calculator) sumArea(shapes ...interface{}) float64 {
	var sum float64
	for _, shape := range shapes {
		switch shape.(type) {
		case circle3:
			r := shape.(circle3).radius
			sum += math.Pi * r * r
		case square3:
			l := shape.(square3).length
			sum += l * l
		}
	}

	return sum
}

func Example3() {
	c := circle3{5}
	s := square3{7}
	calc := calculator{}

	fmt.Printf("total area: %.2f\n", calc.sumArea(c, s))
}
