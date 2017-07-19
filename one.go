package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	type T struct {
		name  string // 对象名
		value int    // 对象值
	}

	fmt.Printf("hello, world\n")
	fmt.Println("The time is", time.Now())
	fmt.Println("number",rand.Intn(10))
	fmt.Println("2222",math.Pi)
}
