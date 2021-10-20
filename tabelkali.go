package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for a := 1; a <= number; a++ {
			fmt.Printf("%d ", i*a)
		}
		fmt.Printf("\n")
	}

}

func main() {

	cetakTablePerkalian(9)
	fmt.Println("Inputkan bilangan tabel perkalian: ")
	var bilangan int
	fmt.Scanf("%d",&bilangan)
	fmt.Printf("Tabel Perkalian Bilangan %d: \n", bilangan)
	cetakTablePerkalian(bilangan)
}