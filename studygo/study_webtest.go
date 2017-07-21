package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayhellowname(w http.ResponseWriter,r *http.Request)  {

	r.ParseForm()//解析参数，默认不解
	fmt.Println(r.Form)//body信息数组
	fmt.Println("url",r.Method)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form {

		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		
	}

	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
func main() {

	http.HandleFunc("/aaa",sayhellowname)//设置路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
