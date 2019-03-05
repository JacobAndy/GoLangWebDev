package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	ghandi := sage{
		Name:  "Ghandi",
		Motto: "Be the change",
	}
	mlk := sage{
		Name:  "Martin Luther King Jr.",
		Motto: "I have a dream",
	}
	jesus := sage{
		Name:  "Jesus",
		Motto: "Let my word be a lamp to your feet",
	}
	sages := []sage{buddha, ghandi, mlk, jesus}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
