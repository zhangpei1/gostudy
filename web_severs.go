package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	t:=time.Now()
	fmt.Fprint(w, "现在时间%d",t)
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:8080", h)
	if err != nil {
		log.Fatal(err)
	}
}

