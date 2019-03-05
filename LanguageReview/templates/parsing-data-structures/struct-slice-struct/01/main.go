package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}
type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	fmt.Println("called")
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	b := sage{
		Name:  "buddha",
		Motto: "Belief in nothing",
	}
	m := sage{
		Name:  "Martin Luther King Jr.",
		Motto: "I have a dream",
	}
	j := sage{
		Name:  "Jesus",
		Motto: "I am coming soon",
	}
	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        4,
	}
	c := car{
		Manufacturer: "Toyota",
		Model:        "Tundra",
		Doors:        4,
	}
	sages := []sage{b, m, j}
	cars := []car{f, c}
	data := items{
		Wisdom:    sages,
		Transport: cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
