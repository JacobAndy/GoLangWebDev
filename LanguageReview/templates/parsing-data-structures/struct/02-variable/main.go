package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type buildC struct {
	Name string
	Age  int
}

func init() {
	fmt.Println("we ran this function")
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	me := buildC{
		Name: "Jacob",
		Age:  18,
	}

	err := tpl.Execute(os.Stdout, me)
	if err != nil {
		log.Fatalln(err)
	}
}
