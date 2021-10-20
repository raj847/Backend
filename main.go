package main

import (
	"fmt"
	"math"
)

func main() {
	//no1
	const phi = 3.14

	fmt.Println("Inpukan T pada tabung:")
	var T float64
	fmt.Scanf("%v \n", &T)
	fmt.Println("Inpukan r pada tabung:")
	var r float64
	fmt.Scanf("%v \n", &r)

	L := 2*phi*r*r + phi*2*r*T
	fmt.Println(L)

	//no2
	var studentScore int = 80

	switch {
	case studentScore > 100 || studentScore < 0:
		fmt.Println("Nilai Invalid")
	case studentScore >= 80:
		fmt.Println("Nilai A")
	case studentScore >= 65:
		fmt.Println("Nilai B")
	case studentScore >= 50:
		fmt.Println("Nilai C")
	case studentScore >= 35:
		fmt.Println("Nilai D")
	case studentScore >= 0:
		fmt.Println("Nilai E")
	}

	//no3
	var bilangan int
	fmt.Println("Inputkan Sebuah Bilangan Bulat : ")
	fmt.Scanf("%d", &bilangan)

	var sqri int = int(math.Sqrt(float64(bilangan)))

	fmt.Println("Faktor dari Bilangan Tersebut Adalah: ")
	var i int
	for i = 1; i <= sqri; i++ {
		if bilangan%i == 0 {
			if bilangan/i == i {
				fmt.Println(i)
			} else {
				fmt.Println(i)
				fmt.Println(bilangan / i)
			}
		}
	}

}