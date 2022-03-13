package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("./hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
		State string
	} {"John Doe", "Awesome"}
	// } {"<script>alert('Howdy!');</script>"}  // unsafe string

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}