package main

import "fmt"
/***
for循环

 */
func main() {
	for_two()
}

//for循环
func for_one() {

	sunm := 0
	for i := 0; i < 10; i++ {
		sunm += i
	}

	fmt.Println(sunm)

}

func for_two() {
	number := 0

	for ; number < 10; number++{

		number += number

		fmt.Println(number)

	}

}
