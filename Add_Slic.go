package main

import "fmt"

func main() {
	var a[]string
	fmt.Printf("%s",a)
	a=append(a, "1","2","3")
	fmt.Printf("a值为==%s,长度为==%d,切片长度为==%v\n",a,len(a),cap(a))
	b:=make([]int,5,5)
	fmt.Printf("b值为==%d,长度为==%d,切片长度为==%v\n",b,len(b),cap(b))

}
