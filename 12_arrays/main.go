package main

import "fmt"

func main() {
	// method 1
	var marks [3]int
	for i := 0; i < 3; i++ {
		marks[i] = (i + 1) * 10
	}
	fmt.Println(marks)

	//method 2
	res := []int{40,50,60,70,80}
	res[4]=90
	// fmt.Println(res)

	res = append(res, marks[:]...)
	fmt.Println(res)
}
