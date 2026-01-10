package main

import "fmt"

func main() {
	var city string
	city = "Chennai"
	var area = "Valsarvakkam" // inferred to string

	// :=
	flat_no :=300
	flat_no = flat_no+1
	pin_code ,plot_no := 600087 ,"77 B/C"
	fmt.Println("City:",city)
	fmt.Println("Area:",area)
	fmt.Println("Flat No:",flat_no)
	fmt.Println("Pin Code:",pin_code)
	fmt.Println("Plot NO:",plot_no)
}
