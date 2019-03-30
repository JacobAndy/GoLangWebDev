package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, req *http.Request) {
	c, e := req.Cookie("session")
	if e != nil {
		id, error := uuid.NewV4()
		if error != nil {
			fmt.Printf("something went wrong: %s", error)
			return
		}
		cookie := &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(c)
}
