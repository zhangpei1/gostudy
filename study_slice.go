package main

import ("fmt"
)

func main() {
	//slice序列的值
	a:=[]int{123132,12,13,13,45,14}
	fmt.Println(len(a))
	fmt.Println(a[:3])
	fmt.Println(a[3:])
	//f := make([]int, 5)
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
