package main

import "fmt"
/****
函数
 */
func main() {
	fmt.Println("x+y",addc(7,9))
	a,b:=addd(10,12)
	fmt.Println("a,b",a, b)
	fmt.Println(split(10))

}
//函数返回值 int 函数参数可以写成 x,y int(如果类型相同)
func addc(x int,y int) int {


	return x+y

}
//返回多个值

func addd(x,y int)  (int,int){

	return x,y

}
//命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

