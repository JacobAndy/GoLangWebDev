package main

import (
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
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	jake := sage{
		Name:  "Jacob",
		Motto: "Work hard, play harder",
	}
	jarid := sage{
		Name:  "Jarid",
		Motto: "hey man",
	}
	joe := sage{
		Name:  "Joe",
		Motto: "Dolla signs",
	}
	aaron := sage{
		Name:  "Aaron",
		Motto: "what up cuh",
	}
	jakeCar := car{
		Manufacturer: "Audi",
		Model:        "rs5",
		Doors:        4,
	}
	jaridCar := car{
		Manufacturer: "Mazda",
		Model:        "3 speed",
		Doors:        4,
	}
	joeCar := car{
		Manufacturer: "Ram",
		Model:        "1500",
		Doors:        4,
	}
	aaronCar := car{
		Manufacturer: "Hyundai",
		Model:        "Sonata",
		Doors:        4,
	}
	sages := []sage{
		jake,
		jarid,
		joe,
		aaron,
	}
	cars := []car{
		jakeCar,
		jaridCar,
		joeCar,
		aaronCar,
	}
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
