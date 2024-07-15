package oops

import "fmt"

type Engine struct {
	Power int
}

func (e Engine) Start() {
	fmt.Println("Engine started with power:", e.Power)
}

type Car struct {
	Brand string
	Engine
}

func Example1() {
	myCar := Car{
		Brand:  "Toyota",
		Engine: Engine{150},
	}

	myCar.Start()
}
