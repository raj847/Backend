package main

import (
	"fmt"
	"strconv"
)

//merubah string jadi integer sumber stackoverflow : https://stackoverflow.com/questions/37765687/golang-how-to-convert-string-to-int
func munculSekali(angka string) []int {
	if _, err := strconv.Atoi(angka); err != nil {
		// do stuff, in case str can not be converted to an int
	}
	var slice []int // empty slice
	for _, digit := range angka {
		slice = append(slice, int(digit)-int('0')) // build up slice
	}
	// return slice
	// 	var slice2 []int = slice[0:1]
	// 	for a := 0; a < len(slice); a++ {
	// 		// ada := false
	// 		for i := 1; i < len(slice2); i++ {
	// 			if slice[a] != slice2[i] {
	// 				slice2 = append(slice2, slice[a])
	// 				// ada = true
	// 			}
	// 		}
	// 		// if ada != false {
	// 		// 	slice2 = append(slice2, slice[a])
	// 		// }
	// 		// fmt.Println(slice)

	// 	}
	// 	return slice2
	// }
	var temp []int
	for i := 0; i < len(slice); i++ {
		beda := false
		for j := 0; j < len(slice); j++ {
			if slice[i] == slice[j] && i != j {
				beda = true
				break
			}
		}

		if !beda {
			temp = append(temp, slice[i])
		}
	}

	return temp
}

func main() {

	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6, 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}
