package main

import "fmt"

func arrayMerge(firstData, secondData []string) []string {
	for i := 0; i < len(secondData); i++ {
		ada := false
		for a := 0; a < len(firstData); a++ {
			if secondData[i] == firstData[a] {
				ada = true
			}
		}
		if !ada {
			firstData = append(firstData, secondData[i])
			// firstData.push(secondData[i])
		}
	}

	return firstData
}

func main() {

	// Test cases

	fmt.Println(arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(arrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(arrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	fmt.Println(arrayMerge([]string{"hwoarang"}, []string{}))

	// ["hwoarang"]

	fmt.Println(arrayMerge([]string{}, []string{}))

	// []

}
