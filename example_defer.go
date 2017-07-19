package main

import "fmt"
/***
defer修饰体

 */
func main() {
	//defer fmt.Println("world")

	//fmt.Println("hello")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}