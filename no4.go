package main

import "fmt"

func sumArray2(arr []int, index1 int, index2 int) int {
	var sum int
	for i := index1; i <= index2; i++ {
		sum += arr[i]
	}
	return sum
}

func MaxSequence(arr []int) int {

	var maxSum int = arr[0]
	var Array2 int

	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			Array2 = sumArray2(arr, i, j)
			if maxSum < Array2 {
				maxSum = Array2
			}
		}
	}

	return maxSum
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}
