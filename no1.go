package main

import (
	"fmt"
	"math"
)

func prima(num int) bool {
	// 	var temp int
	// 	for i := 1; i <= num; i++ {
	// 		if num%i == 0 {
	// 			temp++
	// 		}
	// 	}
	// 	if temp == 2 {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }
	if num < 2 {
		return false
	} else {
		sqrtn := int(math.Sqrt(float64(num)))
		for i := 2; i <= sqrtn; i++ {
			if int(num)%i == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(prima(1000000007)) // true

	fmt.Println(prima(1500450271)) // true

	fmt.Println(prima(1000000000)) // false

	fmt.Println(prima(10000000019)) // true

	fmt.Println(prima(10000000033)) // true
}
