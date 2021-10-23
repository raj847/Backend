package main

import (
	"fmt"
	"math"
)

func Compare(a, b string) string {

	// if len(a) < len(b){
	// 	fmt.Println(a)
	// }
	//  else if len(b) < len(a){
	// 	fmt.Println(b)
	// }
	// else{
	// 	fmt.Println(" ")
	// }
	// if len(a) < len(b) {
	// 	fmt.Println(a)
	// } else if len(b) < len(a) {
	// 	fmt.Println(b)
	// } else {
	// 	fmt.Println(" ")
	// }
	// return Compare(a, b)

	maxLen := int(math.Max(float64(len(a)), float64(len(b))))
	shortLen := int(math.Min(float64(len(a)), float64(len(b))))

	longStr := b
	shortStr := a

	if maxLen == len(a) {
		longStr = a
		shortStr = b
	}
	// if len(a) == len(b) {
	// 	fmt.Println(" ")
	// }

	var compare, stringSama string
	var temp int

	for i := range shortStr {
		for j := range longStr {
			temp = 0
			stringSama = " "

			for int(i+temp) < shortLen && int(j+temp) < maxLen && shortStr[i+temp] == longStr[j+temp] {
				stringSama += string(longStr[j+temp])
				temp += 1
			}

			if len(stringSama) > len(compare) {
				compare = stringSama
			}
		}
	}

	return compare

}

func main() {

	fmt.Println(Compare("AKA", "AKASHI")) // AKA

	fmt.Println(Compare("KANGOORO", "KANG")) // KANG

	fmt.Println(Compare("KI", "KIJANG")) // KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

	fmt.Println(Compare("ILALANG", "ILA")) // ILA

}

// func main() {
// 	str1 := "abc"
// 	str2 := "abd"
// 	fmt.Println(str1 == str2)
// }
