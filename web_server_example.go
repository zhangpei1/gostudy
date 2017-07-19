package main

import (
	"net/http"
	"log"
	"fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

//func (s *Struct) ServeHttp(w http.ResponseWriter,
//	r *http.Request) {
//
//	fmt.Println(s)
//
//}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("PATH: ", r.URL.Path)
	fmt.Println("SCHEME: ", r.URL.Scheme)
	fmt.Println("METHOD: ", r.Method)
	fmt.Println()

	fmt.Fprintf(w, "<h1>Index Page</h1>")
}

func main() {
	http.HandleFunc("/", HandleIndex)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

}
