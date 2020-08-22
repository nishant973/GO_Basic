package main

/* Use append function to dynamicaly grow the slice
   and use[:5] to get a sub-slice , 5/last is exclude and : ---> begin/end */

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]
	arr[0] = 4
	slice[1] = 7
	fmt.Println(arr, slice)

	// Preferred Way

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	slice1 = append(slice1, 9, 7, 56, 48)
	fmt.Println(slice1)

	s1 := slice1[1:]
	s2 := slice1[:2]
	s3 := slice1[1:2]
	fmt.Println(s1, s2, s3)

}
