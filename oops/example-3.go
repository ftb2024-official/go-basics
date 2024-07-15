package oops

import "fmt"

type Printer interface {
	Print()
}

type AdvancedPrinter interface {
	Printer
	Scan()
}

type MFPrinter struct{}

func (mf MFPrinter) Print() {
	fmt.Println("Printing...")
}

func (mf MFPrinter) Scan() {
	fmt.Println("Scanning...")
}

func Example3() {
	mfp := MFPrinter{}

	mfp.Print()
	mfp.Scan()
}
