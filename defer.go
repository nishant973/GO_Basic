package main

import "fmt"

// print 43210
// You can use defecr to some important task like clean up or closing the file
// defer is evaluated immediatley but executed once it's surrounding function is executed.
func main() {
	i := 0
	for ; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}
