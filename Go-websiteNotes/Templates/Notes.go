/*Templates*/

/*Unordered list TODO list*/

/*
data := TodoPageData{
	PageTitle: "My TODO list",
	Todos: []Todo{
		{Title: "Task 1", Done: false},
		{Title: "Task 2", Done: true},
		{Title: "Task 3", Done: true},
	},
}
// to access the data in a template the top most variable is access by {{.}} the dot is a pipeline and the root element of the data.
<h1>{{.PageTitle}}</h1>
<ul>
	{{range .Todos}}
		{{if .Done}}
			<li class="done">{{.Title}}</li?
		{{else}}
			<li {{.Title}}</li>
		{{end}}
	{{end}}
</ul>
*/

/*Control Structures*/
/*
{{/* a comment * /}} 			-> Defines a comment
{{.}} 							-> Renders the root element
{{.Title}} 						-> Renders the "Title"-field in a nested element
{{ if .Done}} {{else}} {{end}}	-> Defines an if-Statment
{{range .Todos}} {{.}} {{end}} 	-> Loops over all "Todos" and renders each using {{.}}
{{block "content" .}} {end}} 	-> Defines a block with the name "content"
*/

/*Parsing Templates from Files*/
// tmpl, err := template.ParseFiles("layout.html")

/*Execute a Template in a Request Handler*/
/*
func(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, "data goes here")
}
*/


/*Main*/

package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO List",
			Todos: []Todo{


				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}


