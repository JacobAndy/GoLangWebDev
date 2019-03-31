package main

import (
	"fmt"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}
func bar(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		fmt.Println("here", un, p, f, l)
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "User name is already taken", http.StatusForbidden)
			return
		}
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		u := user{
			UserName: un,
			Password: bs,
			First:    f,
			Last:     l,
		}
		fmt.Println(un)
		dbUsers[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
