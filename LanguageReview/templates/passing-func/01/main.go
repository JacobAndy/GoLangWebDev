package main

import (
	"log"
	"os"
	"strings"
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

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	jake := sage{
		Name:  "Jacob",
		Motto: "Work hard play hard",
	}
	joe := sage{
		Name:  "Joseph",
		Motto: "Make that money",
	}
	jarid := sage{
		Name:  "jarid",
		Motto: "rona season",
	}
	aaron := sage{
		Name:  "Aaron",
		Motto: "Web Dev",
	}
	jakeCar := car{
		Manufacturer: "Audi",
		Model:        "rs5",
		Doors:        4,
	}
	joeCar := car{
		Manufacturer: "Ram",
		Model:        "1500",
		Doors:        4,
	}
	jaridCar := car{
		Manufacturer: "Mazda",
		Model:        "3 speed",
		Doors:        4,
	}
	aaronCar := car{
		Manufacturer: "Hyundai",
		Model:        "Sonata",
		Doors:        4,
	}
	data := items{
		Wisdom: []sage{
			jake,
			joe,
			jarid,
			aaron,
		},
		Transport: []car{
			jakeCar,
			joeCar,
			jaridCar,
			aaronCar,
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
