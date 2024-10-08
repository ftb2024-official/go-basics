package channelsandgoroutines

import "fmt"

func greet(c chan string) { fmt.Println("Hello, " + <-c + "!") }

func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func Example6() {
	fmt.Println("main() started...")
	cc := make(chan chan string)

	go greeter(cc)

	c := <-cc

	go greet(c)

	c <- "John Doe"
	fmt.Println("main() stopped...")
}
