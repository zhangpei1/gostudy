package main

import (
	"net/http"
	"fmt"
)

type person struct {

}

func main() {
	a:=&person{}
	http.ListenAndServe(":9090",a)

}

func (p *person)ServeHTTP(w http.ResponseWriter,s *http.Request)  {

	s.ParseForm()
	fmt.Fprintln(w,s.Form)

}
