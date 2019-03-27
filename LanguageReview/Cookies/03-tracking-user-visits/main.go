package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", count)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func count(w http.ResponseWriter, req *http.Request) {
	c, e := req.Cookie("views")
	if e == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "views",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)

	io.WriteString(w, c.Value)
}
