package closures

import "fmt"

func createTimer(sec int) func() interface{} {
	timer := func() interface{} {
		if sec > 0 {
			sec--
			return sec
		} else {
			return fmt.Sprintln("Time's up!")
		}
	}

	return timer
}

func Example3() {
	timer := createTimer(3)
	fmt.Println(timer())
	fmt.Println(timer())
	fmt.Println(timer())
	fmt.Println(timer())
	fmt.Println(timer())
	fmt.Println(timer())
	fmt.Println(timer())
	fmt.Println(timer())
}
