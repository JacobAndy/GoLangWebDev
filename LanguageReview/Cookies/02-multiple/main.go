package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "someval",
		Value: "JAKES COOKIE",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}
func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("MY-COOKIE")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}
	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", c2)
	}
	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
	}
}
func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some general value",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some specific value",
	})
	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK YOUR BROWSER")
}
