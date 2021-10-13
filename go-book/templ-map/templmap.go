package main

import (
	"os"
	"text/template"
)

var person = map[string]interface{}{
	"Name":   "John",
	"Age":    24,
	"Gender": "R",
}

var person2 struct {
	Name   string
	Age    int
	Gender string
}

var templ = `----------------------------
Name: 	{{.Name}}
Age: 	{{.Age}}
Gender:	{{.Gender}}
`

func main() {
	person2.Name = "Marry"
	person2.Age = 48
	person2.Gender = "V"

	t := template.Must(template.New("person").Parse(templ))

	_ = t.Execute(os.Stdout, person)
	_ = t.Execute(os.Stdout, person2)
}
