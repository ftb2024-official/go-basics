package strs

import (
	"fmt"
	"strings"
)

func Example2() {
	// проверка наличия подстроки: "strings.Contains(s, substr string) bool": проверяет, содержит ли строка 's' подстроку 'substr'
	fmt.Println(strings.Contains("Hello, world", "world")) // true
	fmt.Println(strings.Contains("Hello, world", "Go"))    // false

	// преобразование регистра: "strings.ToUpper(s string) string": преобразует все символы строки 's' в верхний регистр
	// "strings.ToLower(s string) string": преобразует все символы строки 's' в нижний регистр
	fmt.Println(strings.ToUpper("Hello, world")) // "HELLO, WORLD"
	fmt.Println(strings.ToLower("Hello, world")) // "hello, world"

	// обрезка пробелов и символов: "strings.TrimSpace(s string) string": удаляет пробелы с начала и конца строки 's'
	// "strings.Trim(s, cutset string) string": удаляет все указанные в 'cutset' символы с начала и конца строки 's'
	fmt.Println(strings.TrimSpace("   Hello, world   ")) // "Hello, world"
	fmt.Println(strings.Trim("!!!Hello, world##", "!#")) // "Hello, world"

	// разделение строки: "strings.Split(s, sep string) []string": разделяет строку 's' на подстроки, используя 'sep' как разделитель
	fmt.Println(strings.Split("a,b,c", ",")) // ["a", "b", "c"]

	// объединение строк: "strings.Join(a []string, sep string) string": объединяет элементы массива 'a' в одну строку с разделителем 'sep'
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ", ")) // "a, b, c"

	// замена подстрок: "strings.Replace(s, old, new string, n int) string": заменяет в строке 's' первые 'n' вхождений подстроки 'old' на 'new'. Если
	// 'n' равен '-1', заменяются все вхождения
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // "oinky oinky oink"
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // "moo moo moo"

	// поиск подстроки: "strings.Index(s, substr string) int": возвращает индекс первого вхождения подстроки 'substr' в строке 's' или '-1', если
	// подстрока не найдена
	// "strings.LastIndex(s, substr string) int": возвращает индекс последнего вхождения подстроки 'substr' в строке 's' или '-1', если
	// подстрока не найдена
	fmt.Println(strings.Index("Hello, world", "world"))            // 7
	fmt.Println(strings.LastIndex("Hello, world, world", "world")) // 14
	fmt.Println(strings.Index("Hello, world", "Go"))               // -1
}
