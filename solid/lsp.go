package solid

import "fmt"

type person interface {
	getName() string
}

type printer struct{}

func (printer) info(p person) { fmt.Println("Name:", p.getName()) }

type human struct{ name string }

func (h human) getName() string { return h.name }

type student struct {
	human
	grades map[string]int
}

type teacher struct {
	human
	degree string
	salary float64
}

func Example5() {
	h := human{"Alex"}
	t := teacher{human{"Michael"}, "CS", 2000}
	s := student{human{"Mike"}, map[string]int{"English": 8, "Math": 9}}
	var p printer

	p.info(h)
	p.info(t)
	p.info(s)
}
