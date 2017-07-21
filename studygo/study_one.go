package main

import (
	"fmt"
)

func main() {
	b := getdata()
	for i := 0; i < 10; i++ {
		//fmt.Println("a=",a(1))
		fmt.Println("b=", b(2))

	}

}
func getdata() func(int) int {

	sum := 0
	fmt.Println("sum=", sum)
	return func(x int) int {

		sum += x
		return sum
	}

}
