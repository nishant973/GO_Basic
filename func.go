package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	iStarted, err := startWebServer(port, 3)
	fmt.Println(iStarted, err)
}

/* Though it's illogical to return both , but it's just for demo as you can implement dynamic return type */
func startWebServer(port, retriesCount int) (bool, error) {
	fmt.Println("Server started to retry : ", retriesCount, "times")
	fmt.Println("Server started at port : ", port)
	return true, errors.New("Something went wrong")
}
