package main

import "fmt"

func BinarySearch(array []int, x int) {

	a, b := 0, len(array)-1

	if x > array[b] || x < array[0] {
		fmt.Println(-1)
		return
	} else if array[a] == x {
		fmt.Println(a)
		return
	} else if array[b] == x {
		fmt.Println(b)
		return
	} else {
		for a+1 != b {
			if array[(a+b)/2] == x {
				fmt.Println((a + b) / 2)
				return
			} else if array[(a+b/2)] < x {
				a = (a + b) / 2
			} else {
				b = (a + b) / 2
			}
		}
	}

	fmt.Println(-1)
	return
}

func main() {

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1
}
