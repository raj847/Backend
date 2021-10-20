	//no5
	package main

	import "fmt"

	func palindrome(kata string)bool {
		for i := 0; i < len(kata)/2; i++ {
			if kata[i] != kata[len(kata)-1-i] {
			  return false;
			}
		  }
		  return true;
		}

	// 	var p string
	// 	for i:= len(kata)-1 ; i>=0;i--{
	// 		p+=kata[i]
	// 	}

	// 	if p==kata{
	// 		return true
	// 	}else{
	// 		return false
	// 	}

	// }

		func main() {
			fmt.Println(palindrome("civic")); 
			fmt.Println(palindrome("katak")); 
			fmt.Println(palindrome("kasur rusak")); 
			fmt.Println(palindrome("mistar")); 
			fmt.Println(palindrome("lion")); 
		}
		
		