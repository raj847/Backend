package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	sMap := make(map[string]int)
	for _, i := range s.score {
		sMap[""] += i
	}

	return float64(sMap[""] / len(s.score))
}

func (s Student) Min() (min int, name string) {
	sMap := map[string]int{
		"": s.score[0],
	}

	var temp int

	for i, j := range s.score {
		if sMap[""] > j {
			sMap[""] = j
			temp = i
		}
	}

	return sMap[""], s.name[temp]
}

func (s Student) Max() (max int, name string) {
	sMap := map[string]int{
		"": s.score[0],
	}

	var temp int

	for i, j := range s.score {
		if sMap[""] < j {
			sMap[""] = j
			temp = i
		}
	}

	return sMap[""], s.name[temp]
}

func main() {

	var a = Student{}

	for i := 0; i < 6; i++ {

		var name string

		fmt.Print("Input " + string(i) + " Student’s Name : ")

		fmt.Scan(&name)

		a.name = append(a.name, name)

		var score int

		fmt.Print("Input " + name + " Score : ")

		fmt.Scan(&score)

		a.score = append(a.score, score)

	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())

	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

}
