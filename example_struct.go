package main

import "fmt"
/***
结构体

 */

type Fare struct {
	X int
	Y int
}

var (
	v1 = Fare{1, 2}  // 类型为 Vertex
	v2 = Fare{X: 1}  // Y:0 被省略
	v3 = Fare{}      // X:0 和 Y:0
	p  = &Fare{1, 2} // 类型为 *Vertex
)

func main() {
	v:=Fare{6,7}
	v.X=7
	if v.X==5 {

		println(v.X)

	}
	b:=&v
	b.X=1
	fmt.Println(v)
	fmt.Println(v1,v2,v3,*p)
}
