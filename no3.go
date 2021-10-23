package main

import "fmt"

func swap(a, b *int) {
	tempa := *a
	tempb := *b
	*a = tempb
	*b = tempa
}
func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("a sebelum swap ", a)
	fmt.Println("b sebelum swap ", b)
	swap(&a, &b)

	fmt.Println("a sesudah swap", a)
	fmt.Println("b sesudah swap", b)
}
