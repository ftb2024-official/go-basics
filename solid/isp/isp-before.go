package isp

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	volume() float64
}

type square struct{ length float64 }
type cube struct{ length float64 }

func (s square) area() float64   { return s.length * s.length }
func (s square) volume() float64 { return 0 } // square is flat surface -> no volume

func (c cube) area() float64   { return math.Pow(c.length, 2) }
func (c cube) volume() float64 { return math.Pow(c.length, 3) }

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, sh := range shapes {
		sum += sh.area()
	}
	return sum
}

func areaVolumeSum(shapes ...shape) float64 {
	var sum float64
	for _, sh := range shapes {
		sum += sh.area() + sh.volume()
	}

	return sum
}

func Example1() {
	s := square{3}
	c := cube{4}

	fmt.Println(areaSum(s, c))
	fmt.Println(areaVolumeSum(s, c))
}
