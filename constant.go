package main

// Constant and iota basic

import "fmt"

const (
	first  = iota
	second = iota
	third  = iota + 6
	fourth
)

const (
	refirst  = iota
	resecond = iota
)

func main() {
	//const first = 345 // allows to insitialize again
	const a = 123
	// cons b int = 7
	fmt.Println(a)
	fmt.Println(first, second, third, fourth)
	fmt.Println(refirst, resecond)
	fmt.Println(first)
	fmt.Println(a + 5)
	fmt.Println(a + 6.2)
	// fmt.Println((float32)b + 6.2) // not allowed
}
