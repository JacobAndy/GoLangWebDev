package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	bot := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}
	err := tpl.Execute(os.Stdout, bot)
	if err != nil {
		log.Fatalln(err)
	}
}
