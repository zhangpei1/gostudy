package main

import "fmt"

var a  =[]int{1,2,3,10,5,6,7,8}
func main() {
	for i, v := range a {
		fmt.Printf("2**%d = %d\n", i, v)

	}
	for _, value := range a {
		fmt.Printf("%d\n", value)
	}
	}
