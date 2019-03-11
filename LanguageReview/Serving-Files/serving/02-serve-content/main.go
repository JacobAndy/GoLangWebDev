package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", jake)
	http.HandleFunc("/toby.jpg", jakeImg)
	http.ListenAndServe(":8080", nil)
}

func jake(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="toby.jpg">`)
}

func jakeImg(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("beauty.jpeg")
	if err != nil {
		http.Error(res, "file not found", 404)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}
