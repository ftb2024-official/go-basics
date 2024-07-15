package oops

import "fmt"

type Drivable interface {
	Drive()
}

type Vehicle struct {
	Brand string
}

func (v Vehicle) Drive() {
	fmt.Println(v.Brand, "is driving")
}

func startDriving(d Drivable) {
	d.Drive()
}

func Example2() {
	bmw := Vehicle{Brand: "BMW"}
	bike := Vehicle{"Yamaha"}

	startDriving(bmw)
	startDriving(bike)
}
