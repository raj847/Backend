// no4
package main

import "fmt"

func primeNumber(number int) bool {
	var p int
	for i := 1; i <= number; i++{
		if number%i == 0 {
			p+=1
		}
	}
		if p==2{
			return true
		}else{
			return false
		}
	}


func main() {
	fmt.Println(primeNumber(11)) // true
	fmt.Println(primeNumber(13)) // true
	fmt.Println(primeNumber(17)) // true
	fmt.Println(primeNumber(20)) // false
	fmt.Println(primeNumber(35)) // false
	fmt.Println("Inputkan Sebuah Bilangan: ")
	var number int
	fmt.Scanf("%d",&number)
	if primeNumber(number) == true {
		fmt.Println("Termasuk Bilangan Prima")
	} else {
		fmt.Println("Bukan Termasuk Bilangan Prima")
	}
}