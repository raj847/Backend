//no6
package main

import "fmt"

func pangkat(base, pangkat int) int {
	var total int = 1
	for i := 0; i < pangkat; i++ {
		total *= base
	}
	return total
}

func main() {
	fmt.Println(pangkat(2, 3))  // 8
	fmt.Println(pangkat(5, 3))  // 125
	fmt.Println(pangkat(10, 2)) // 100
	fmt.Println(pangkat(2, 5))  // 32
	fmt.Println(pangkat(7, 3))  // 343
	fmt.Println("Masukkan Nilai x : ")
	var x int
	fmt.Scanf("%v \n", &x)
	fmt.Println("Masukkan Nilai n : ")
	var n int
	fmt.Scanf("%v \n", &n)
	fmt.Println("total dari x pangkat n adalah ", pangkat(x, n))
}