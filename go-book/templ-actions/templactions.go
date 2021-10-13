package main

import (
	"fmt"
	"os"
	"text/template"
)

var templ = `{{/* Testing different kinds of actions */}}---------------------
Name: 			{{.Name}}
Citizen? 		{{if .Citizen}}Yes{{else}}No{{end}}
Resident? 		{{if .Resident}}Yes{{else}}No{{end}}
Setup?			{{if and .Registered .Employed}}Yes{{else}}No{{end}}
Legal?			{{if or .Citizen .Resident}}Yes{{else}}No{{end}}
Address:		{{(index .Contacts 2).Value | printf "%s"}}

Contacts ({{len .Contacts}}):	{{template "contacts" .Contacts}}
{{.Tag | html}}
---------------------{{define "contacts"}}
{{range .}}  - {{.Type}}:	{{.Value}}
{{end}}{{end}}`

type Person struct {
	Name                                    string
	Citizen, Resident, Registered, Employed bool
	Contacts                                []*Contact
	Document                                string
}

type Contact struct {
	Type  string
	Value string
}

func (p Person) Tag() string {
	return fmt.Sprintf("<div><h1>%s</h1><span>%s</span></div>", p.Name, p.Document)
}

func main() {
	t := template.Must(template.
		New("actions").
		Parse(templ))

	p := Person{
		Name:       "John Doe",
		Citizen:    false,
		Resident:   true,
		Registered: true,
		Employed:   true,
		Contacts: []*Contact{
			{
				Type:  "Phone",
				Value: "+31987654321",
			},
			{
				Type:  "Email",
				Value: "john@doe.com",
			},
			{
				Type:  "Address",
				Value: "Jonstraat, 999",
			},
		},
		Document: "9876543212345678-9",
	}

	_ = t.Execute(os.Stdout, p)
}
