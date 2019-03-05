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
	sages := map[string]string{
		"Football": "Tim Tebow",
		"UFC":      "Khabib",
		"Soccer":   "Some guy from the EU",
		"Cricket":  "lol",
		"theBest":  "ME",
	}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
