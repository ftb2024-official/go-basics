package defers

import (
	"fmt"
	"io"
	"os"
)

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close() // гарантирует, что файл будет закрыт при выходе из функции

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Example1() {
	d, err := readFile("./test.txt")
	if err != nil {
		fmt.Sprintln("Can't find test file")
	}

	fmt.Println(d)
}
