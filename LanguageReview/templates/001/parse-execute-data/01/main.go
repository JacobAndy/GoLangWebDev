package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	/*
		we pass data into a file as the 3rd parameter of Executing a template,
		inside the file, we refer to that data being passed in as {{.}}

			ONLY 1 PIECE OF DATA IS ALLOWED TO BE PASSED
	*/
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
