package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hot dog, Hot dog, Hot diggity DOG")
}
func (h hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hot Cat, Hot Cat, Hot Diggity Cat!")
}

func main() {
	var d hotdog
	var c hotcat
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)

	http.ListenAndServe(":8080", mux)
}
