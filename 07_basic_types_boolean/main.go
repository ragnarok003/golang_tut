package main

import "fmt"

func main(){
	ghostMode:= true
	killMode:= false

	fmt.Println("Stealth Mode:",(ghostMode && killMode))
	fmt.Println("Rage Mode:",(ghostMode || killMode))
}
