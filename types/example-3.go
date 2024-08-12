package types

import (
	"fmt"
)

func describe(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Interger:", v)
	case float32:
		fmt.Println("Float:", v)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	case byte:
		fmt.Println("Byte:", v)
	case rune:
		fmt.Println("Rune:", v)
	case []interface{}:
		fmt.Println("Slice:", v)
	case map[interface{}]interface{}:
		fmt.Println("Map:", v)
	default:
		fmt.Println("Unknown type:", v)
	}
}

func Example3() {
	describe(42)
	describe("halo")
	describe(true)
	describe(byte('a'))
	describe(rune('z'))
	describe([]interface{}{1, 2, 3, 4, 5})
	describe(map[interface{}]interface{}{1: "one"})
	describe(3.14)
}
