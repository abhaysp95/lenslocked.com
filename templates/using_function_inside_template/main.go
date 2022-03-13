package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Admin bool
	ID int
	Email string
	// HasPermission func(string) bool
}

type ViewData struct {
	User
	Usage, Limit, Warning int
}

var testTemplate *template.Template

/* func (u User) HasPermission(feature string) bool {
	if feature == "feature-a" {
		return true
	} else {
		return false
	}
} */

func main() {
	// testTemplate = template.Must(template.ParseFiles("./hello.gohtml"))
	testTemplate = template.Must(template.New("hello.gohtml").Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			if feature == "feature-a" {
				return true
			} else {
				return false
			}
		},
		"ifIE": func() template.HTML {
			return template.HTML("<!--[if IE]>")
		},
		"endif": func() template.HTML {
			return template.HTML("<![endif]-->")
		},
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	}).ParseFiles("./hello.gohtml"))

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{
		User{
			Admin: true,
			ID: 1,
			Email: "someone@somewhere.xyz",
		},
		16, 12, 15,
	}
	// err := testTemplate.Execute(w, vd)
	// creating and clone and passing a closure
	err := template.Must(testTemplate.Clone()).Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			// or could have user variable of User type (also passed in ViewData)
			if vd.ID == 1 && feature == "feature-a" {
				return true
			} else {
				return false
			}
		},
	}).Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
