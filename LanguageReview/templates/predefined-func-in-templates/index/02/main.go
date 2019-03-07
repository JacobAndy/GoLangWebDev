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
	xs := []string{
		"Jake",
		"Joe",
		"Aaron",
		"Jarid",
	}
	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Anderson",
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
