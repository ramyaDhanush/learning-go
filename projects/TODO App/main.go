/*

ServeMux - HTTP request multiplexer
[Matches the URL of each incoming request against a
list of registered patterns and calls the handler]

*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}


func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: true},
			{Item: "Build Projects Go", Done: false},
		},
	}
	
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)
	
	log.Fatal(http.ListenAndServe(":9090", mux))
}