package main

import (
	"fmt"
	"math"
)

func primaSegiEmpat(high, wide, start int) {
	fmt.Printf("\n\n")

	var sum int
	// your code here
	mapBilPrima := mapBilPrima(high, wide, start)
	j := 1
	for i := 1; i <= wide; i++ {
		for j <= i*high {
			fmt.Printf("%d\t", mapBilPrima[j])
			j++
		}
		fmt.Printf("\n\n\n")
	}

	for _, v := range mapBilPrima {
		sum += v
	}

	println(sum)

	fmt.Printf("\n")
}

func mapBilPrima(high, wide, start int) map[int]int {

	var index int
	indexBilPrima := make(map[int]int)

	a := start + 1

	for a >= start+1 {
		isPrime := true
		sqrtn := int(math.Sqrt(float64(a)))
		b := 2
		for b <= sqrtn {
			if a%b == 0 {
				isPrime = false
				b = a
			}
			b++
		}

		if isPrime {
			index++
			indexBilPrima[index] = a
		}

		if index == high*wide {
			break
		}

		a++
	}

	return indexBilPrima
}

func main() {

	primaSegiEmpat(2, 3, 13)

	primaSegiEmpat(5, 2, 1)

}
