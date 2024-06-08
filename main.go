package main

import (
	"html/template"
	"net/http"
	"slices"
)

type Todo struct {
	Title    string
	Progress bool //if false it means not started
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var progressIndexEnd int = 0 //the last index with an in progress element

func AddItem(list *[]Todo, item Todo) {
	if item.Progress {
		slices.Insert(list, progressIndexEnd, item)
		progressIndexEnd++
	} else {
		*list = append(*list, item)
	}
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos:     []Todo{
				//{Title: "Task 1", Done: false},
				//{Title: "Task 2", Done: true},
				//{Title: "Task 3", Done: true},
			},
		}

		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
