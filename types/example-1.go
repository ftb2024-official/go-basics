package types

import (
	"fmt"
	"reflect"
)

func Example1() {
	var x interface{} = 42

	fmt.Println("Type:", reflect.TypeOf(x))
	fmt.Println("Value:", reflect.ValueOf(x))
}
