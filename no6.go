package main

import (
	"fmt"
	"strconv"
)

func FindMax(arr []int) int {

	// your code here
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max

}

func MaximumBuyProduct(money int, productPrice []int) {

	// your code here
	var maxProductTerbeli []int

	for i := 0; i < len(productPrice); i++ {
		terbeli := productPrice[i]
		productTerbeli := 1
		// fmt.Println(terbeli)

		if terbeli > money {
			terbeli -= productPrice[i]
			productTerbeli--
		} else {
			for j := i + 1; j < len(productPrice); j++ {
				terbeli += productPrice[j]
				productTerbeli++
				if terbeli > money {
					terbeli -= productPrice[j]
					productTerbeli--
				}
			}
			maxProductTerbeli = append(maxProductTerbeli, productTerbeli)
		}
	}

	fmt.Println(strconv.Itoa(FindMax(maxProductTerbeli)))
}

func main() {

	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}) // 3

	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4

	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}) // 4

	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}) // 1

	MaximumBuyProduct(0, []int{10000, 30000}) // 0

}
