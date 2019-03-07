package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	jake := user{
		"Jacob",
		"Work hard play hard",
		true,
	}
	joe := user{
		"Joseph",
		"Make that money",
		false,
	}
	aaron := user{
		"Aaron Blackshear",
		"say what!?",
		true,
	}
	jarid := user{
		"Jarid",
		"Some random text in here",
		false,
	}
	noBody := user{
		"",
		"nobody",
		false,
	}
	users := []user{
		jake,
		joe,
		aaron,
		jarid,
		noBody,
	}
	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
