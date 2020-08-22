package main

import "fmt"

func main() {
	m := map[string]int{"nishi": 27, "akki": 43}
	fmt.Println(m)
	m["nishi"] = 65
	fmt.Println(m["nishi"], m["ni"])
	delete(m, "nishi")
	fmt.Println(m)
}
