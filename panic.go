package main

import "fmt"

// Just a basic overview of panic , just like exception in other language
func main() {
	fmt.Println("Server Started")

	panic("Somrthing bad happenend")

	//fmt.Println("logic implemeted !!")
}
