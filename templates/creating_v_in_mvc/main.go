package main

import (
	"net/http"

	"lenslocked.com/templates/creating_v_in_mvc/view"
)

var index *view.View
var contact *view.View

func main() {
	index = view.NewView("bootstrap", "view/index.gohtml")
	contact = view.NewView("bootstrap", "view/contact.gohtml")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	err := index.Render(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contactHandler(w http.ResponseWriter, req *http.Request) {
	err := contact.Render(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
