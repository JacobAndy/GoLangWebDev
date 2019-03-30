package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<a href="/set"><h1>Index page, click here to set your cookie</h1></a>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "jacob_is_awesome",
		Value: "yes, very true",
	})
	fmt.Fprintln(w, `<a href="/read"><h1>We set your cookie, click here to read your cookie</h1></a>`)
}
func read(w http.ResponseWriter, req *http.Request) {
	c, e := req.Cookie("jacob_is_awesome")
	if e != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintln(w, `<a href="/expire"><h1>Your cookie: %v </h1></a>`, c)
}
func expire(w http.ResponseWriter, req *http.Request) {
	c, e := req.Cookie("jacob_is_awesome")
	if e != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
