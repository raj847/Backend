package main

import "fmt"

func pow(x, n int) int {

	result := 1
	if n == 0 {
		return result
	}

	for n >= 1 {
		if n%2 != 0 {
			result *= x
		}

		x *= x
		n /= 2
	}

	return result

}

func main() {

	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(7, 2)) // 49

	fmt.Println(pow(10, 5)) // 100000

	fmt.Println(pow(17, 6)) // 24137569

	fmt.Println(pow(5, 3)) // 125

}
