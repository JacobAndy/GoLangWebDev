package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", jake)
	http.HandleFunc("/toby.jpg", jakeImg)
	http.ListenAndServe(":8080", nil)
}
func jake(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<img src="toby.jpg">
	`)
}
func jakeImg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "beauty.jpeg")
}
