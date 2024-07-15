package functions

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func Example3() {
	sayHello := greet

	sayHello("Alice")
	sayHello("Bob")
}
