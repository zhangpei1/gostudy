package main

import "fmt"
var c, python, java bool=true,true,true
var (b bool=false
	MaxInt uint64     = 1<<64 - 1
    n int8=12
)

func main() {
	var i int
	//只能在函数体内使用，函数体外必须用var
	k:=8
	fmt.Println(i, c, python, java,k)
	const f = "%T(%v)\n"
	fmt.Printf(f, b,b)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, n,n)

}