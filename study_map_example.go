package main

import "fmt"

type kkk struct {

	a int
	b int


}

func main() {
	var a map[string]kkk
	a=make(map[string]kkk)
	a["dsa"]=kkk{1,1}
	fmt.Println(a)
}
