package main

import (
	"fmt"
	"strings"
)

//sumber https://stackoverflow.com/questions/19310700/what-is-a-rune

func caesar(offset int, input string) string {
	transform := func(r rune) rune {
		if offset > 26 {
			offset = offset % 26
		}

		i := int(r) + offset
		if i > 'z' {
			return rune(i - 26)
		} else if i < 'a' {
			return rune(i + 26)
		}
		return rune(i)
	}

	result := strings.Map(transform, input)

	return result
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}
