package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type hotdog int

func (h hotdog) ServeHTTP(r http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	tpl.ExecuteTemplate(r, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

var tpl *template.Template

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
