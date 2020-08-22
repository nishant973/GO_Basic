package main

import "fmt"

// Pointer basic

func main() {
	firstName := "Nishant"
	var ptr *string = &firstName

	var lastName *string = new(string)
	*lastName = "J"
	fmt.Println(ptr, *ptr)
	fmt.Println(lastName)
	fmt.Println(*lastName)
}
