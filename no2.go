package main

import "fmt"

func moneyCoins(money int) []int {

	pecahanUang := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	hasil := []int{}
	// your code here
	// fmt.Println(len(pecahanUang))
	for i := len(pecahanUang) - 1; i >= 0; i-- {
		// fmt.Println(len(pecahanUang) - 1)
		for pecahanUang[i] <= money {
			hasil = append(hasil, pecahanUang[i])
			money -= pecahanUang[i]
			if money == 0 {
				return hasil
			}
		}
	}
	return hasil
}

func main() {

	fmt.Println(moneyCoins(123)) // [100 20 1 1 1]

	fmt.Println(moneyCoins(432)) // [200 200 20 10 1 1]

	fmt.Println(moneyCoins(543)) // [500, 20, 20, 1, 1, 1]

	fmt.Println(moneyCoins(7752)) // [5000, 2000, 500, 200, 50, 1, 1]

	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]

}
