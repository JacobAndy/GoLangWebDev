package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/careers":
		io.WriteString(res, "Join our career team today!")
	case "/about":
		io.WriteString(res, "Learn more about us!")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
