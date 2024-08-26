package isp

import (
	"fmt"
	"math"
)

type shape2 interface {
	area() float64
}

type object2 interface {
	shape2
	volume() float64
}

type square2 struct{ length float64 }
type cube2 struct{ square2 }

func (s square2) area() float64 { return s.length * s.length }
func (c cube2) volume() float64 { return math.Pow(c.length, 3) }

func areaSum2(shapes ...shape2) float64 {
	var sum float64
	for _, sh := range shapes {
		sum += sh.area()
	}
	return sum
}

func areaVolumeSum2(shapes ...object2) float64 {
	var sum float64
	for _, sh := range shapes {
		sum += sh.area() + sh.volume()
	}
	return sum
}

func Example2() {
	s1 := square2{5}
	s2 := square2{6}
	c1 := cube2{square2{3}}
	c2 := cube2{square2{4}}

	fmt.Println(areaSum2(s1, s2, c1, c2))
	fmt.Println(areaVolumeSum2(c1, c2))
}
