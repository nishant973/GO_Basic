package main

import "fmt"

func main() {
	type user struct {
		ID        int
		firstName string
		lastName  string
	}
	var u user
	u.ID = 1
	u.firstName = "Nishant"
	u.lastName = "J"
	fmt.Println(u)

	// other way
	u2 := user{ID: 2,
		firstName: "akki",
		lastName:  "Jay",
	}
	fmt.Println(u2)

}
