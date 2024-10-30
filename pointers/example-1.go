package pointers

import "fmt"

func Example1() {
	a := 2
	fmt.Println("a(7) =", a)
	b := &a
	fmt.Println("\nb(9) =", b)
	fmt.Println("*b(10) =", *b)

	*b = 3
	fmt.Println("\na(13) =", a)
	fmt.Println("b(14) =", b)
	fmt.Println("*b(15) =", *b)
	c := &a
	fmt.Println("\nc(17) =", c)

	d := new(int)
	fmt.Println("\n*d(20) =", *d)
	fmt.Println("d(21) =", d)

	*d = 12
	fmt.Println("*d(24) =", *d)

	*c = *d
	fmt.Println("\na(27) =", a)
	fmt.Println("*c(28) =", *c)
	fmt.Println("c(29) =", c)

	*d = 13
	fmt.Println("\n*d(32) =", *d)
	fmt.Println("d(33) =", d)

	fmt.Println("\na(35) =", a)
	fmt.Println("*c(36) =", *c)
	fmt.Println("c(37) =", c)

	c = d
	*c = 100
	fmt.Println("\na(41) =", a)
	fmt.Println("*c(42) =", *c)
	fmt.Println("c(43) =", c)

	fmt.Println("\n*d(45) =", *d)
	fmt.Println("d(46) =", d)
}
