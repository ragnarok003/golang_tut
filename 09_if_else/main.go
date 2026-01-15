package main

import "fmt"

func main(){
	score:=72.0

	if percent:=(score/10); percent>=9.0{ // short hand
		fmt.Println("A",percent)
	}else if (score/10) >=7.5{
		fmt.Println("B",percent)
	}else if (score/10) >=4.5{
		fmt.Println("C",percent)
	}else{
		fmt.Println("D",percent)
	}

	if score<4.5{ //normal if statement 
		fmt.Println("Fail")
	}else{
		fmt.Println("Pass")
	}
}
