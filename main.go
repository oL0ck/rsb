package main

import (
	"fmt"
	"net/http"
	"time" //some workaround
)

var stop bool

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Hello, World!",time.Now())
	//stop = true
}
func main(){
	stop = false
	http.HandleFunc("/", HomeRouterHandler)
	go http.ListenAndServe(":8090", nil)
	for !stop {
	}
}

