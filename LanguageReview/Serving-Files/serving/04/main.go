package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", jake)
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
func jake(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/img/beauty.jpeg">`)
}
