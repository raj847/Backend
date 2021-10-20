package main

import "fmt"

func playWithAsterik(n int) {

	for i := n - 1; i >= 0; i-- {
		for a := 0; a < i; a++ {
			fmt.Printf(" ")
		}
		for b := 0; b < n-i; b++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func main() {

	playWithAsterik(5)
	fmt.Println("Masukkan Jumlah Bintang: ")
	var jumlah int
	fmt.Scan(&jumlah)
	playWithAsterik(jumlah)
}