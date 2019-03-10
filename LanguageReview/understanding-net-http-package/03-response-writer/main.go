package main

import (
	"fmt"
	"net/http"
)

type hotdog int

////////////////////////////response				//request
func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Anderson-key", "this is from Jacob")
	res.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(res, "<h1>Any code you want, can go in here</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
