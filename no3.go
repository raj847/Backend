package main

import "fmt"

func hapus(slice []int, s int) {
	slice = append(slice[:s], slice[s+1:]...)
}

func DragonOfLoowater(dragonHead, knightHeight []int) {

	// your code here
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("knight fall")
		return
	}

	var minimumTotalHeight, tempH int
	var isDie bool
	var index int

	for _, diameter := range dragonHead {
		tempH = 0
		isDie = false
		for i, height := range knightHeight {
			if height >= diameter && height < tempH {
				minimumTotalHeight += height - tempH
				tempH = height
				index = i
				continue
			} else if isDie {
				continue
			} else if height >= diameter {
				minimumTotalHeight += height
				tempH = height
				index = i
				isDie = true
			}
		}

		if !isDie {
			fmt.Println("knight fall")
			return
		}

		hapus(knightHeight, index)
	}

	fmt.Println(minimumTotalHeight)
	return
}

func main() {

	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}
