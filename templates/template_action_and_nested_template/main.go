package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

type Widget struct {
	Name string
	Price int
}

type ViewData struct {
	Name string
	Widgets []Widget
}

func action_and_nesting() {
	// in case of error, below line panics, because of Must()
	testTemplate = template.Must(template.ParseFiles("hello.gohtml"))

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	vd := ViewData{
		Name: "John Doe",
		Widgets: []Widget{
			{"Red", 12},
			{"Green", 12},
			{"Blue", 12},
		},
	}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	action_and_nesting()
}
