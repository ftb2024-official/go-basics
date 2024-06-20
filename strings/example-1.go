package strings

import "fmt"

func Example1() {
	str := "Привет, мир!"

	fmt.Println(str[0]) // доступ к символу (на самом деле к байту)

	// работа с рунами для доступа к символам Unicode
	for _, runeVal := range str {
		fmt.Printf("%#U ", runeVal)
	}
}
