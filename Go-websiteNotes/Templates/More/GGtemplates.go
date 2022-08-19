package main

import (
	"os"
	"fmt"
	"text/template"
)


type Sesh struct {

	Authenticated bool
	User string
	Power SuperPower
}

type SuperPower struct{
	PowerID int64
	Name string
	ability string

}

func main() {

	obj1 := SuperPower{01,"Flying", "mobility"}
	obj2 := Sesh{true, "hunter", obj1}

	tmpl := template.New("Template1")

	tmpl, _ = tmpl.Parse("Hello {{.User}}, your super power is {{.Power.Name}}!")

	err := tmpl.Execute(os.Stdout, obj2)

	if err != nil {
		fmt.Println(err)
	}
}
