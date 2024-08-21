package solid

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type shape interface {
	name() string
	area() float64
}

type outputter struct{}

func (o outputter) Text(sh shape) string {
	return fmt.Sprintf("area of %s: %.2f", sh.name(), sh.area())
}

func (o outputter) JSON(sh shape) string {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: sh.name(),
		Area: sh.area(),
	}

	bRes, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	return string(bRes)
}

type circle2 struct{ radius float64 }

func (c circle2) name() string  { return "circle" }
func (c circle2) area() float64 { return math.Pi * c.radius * c.radius }

type square2 struct{ length float64 }

func (s square2) name() string  { return "square" }
func (s square2) area() float64 { return s.length * s.length }

func Example2() {
	var out outputter
	c := circle2{5}
	s := square2{7}

	fmt.Println(out.Text(c))
	fmt.Println(out.JSON(c))
	fmt.Println(out.Text(s))
	fmt.Println(out.JSON(s))
}
