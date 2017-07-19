package main

import (
	"fmt"
	"math"
)

type A struct {
	b, c float32
}

var m map[int]A
var t = map[string]A{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[int]A)
	m[0] = A{
		40.68433, -74.39967,
	}
	fmt.Println(m[0], len(m))

	t["Bell Labs"] = A{123132, 20.12312}
	//删除元素：
	delete(t, "Bell Labs")
	fmt.Println(t)
	//通过双赋值检测某个键存在：
	v, ok := t["Google"]
	fmt.Println("The value:", v, "Present?", ok)
	//函数值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(i),
		)
	}
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
