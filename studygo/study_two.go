package main

import "fmt"

func main() {
	h:=HuanMa{"小杨","dsad"}
	b:=BaoMa{h,8,"女"}
	fmt.Println(b.Kaiche())
	
}
type HuanMa struct {
	 ss string
	 name string
}

type BaoMa struct {
	HuanMa
	age int8
	sex string
}

func (d *BaoMa) Kaiche() string  {


	return d.name+d.sex
	
}
