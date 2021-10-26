package main

import (
	"fmt"
)

func sumEl(arr []int) int {

	// your code here
	var sum int
	for _, i := range arr {
		sum += i
	}

	return sum
}

func canOut(inHand, deck []int) bool {
	if len(inHand) != len(deck) {
		return false
	}

	for _, j := range inHand {
		for _, k := range deck {
			if j == k {
				return true
			}
		}
	}

	return false
}

func playingDomino(cards [][]int, deck []int) interface{} {

	// your code here
	var mengeluarkanKartu []int

	for i, j := range cards {
		if canOut(j, deck) {
			mengeluarkanKartu = append(mengeluarkanKartu, i)
		}
	}

	if len(mengeluarkanKartu) == 0 {
		return "tutup kartu"
	}

	kartuKeluar := cards[mengeluarkanKartu[0]]

	for _, j := range mengeluarkanKartu {
		if sumEl(cards[j]) > sumEl(kartuKeluar) {
			kartuKeluar = cards[j]
		}
	}

	return kartuKeluar
}

func main() {

	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))

	// [3, 4]

	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))

	// [6 5]

	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))

	// “tutup kartu”

}
