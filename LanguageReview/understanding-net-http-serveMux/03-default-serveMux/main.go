package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hot dog hot dog hot diggity dog")
}
func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hot cat hot cat hot diggity cat")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)
	http.ListenAndServe(":8080", nil)
}
