package main

import (
	"html/template"
	"os"
)

type Test struct {
	HTML string
	SafeHTML template.HTML  // don't use the type when dealing with user input (code injection)
	Title string
	Path string
	Dog Dog
	Map map[string]string
}

type Dog struct {
	Name string
	Age int
}

func context_encoding() {
	t, err := template.ParseFiles("./context.gohtml")
	if err != nil {
		panic(err)
	}

	data := Test {
		HTML: "<h1>A header!</h1>",
		SafeHTML: template.HTML("<h1>A safe header</h1>"),
		Title: "Backslash! An in depth look at \"\\\" character",
		Path: "/dashboard/settings",
		Dog: Dog {"Fido", 6},  // structs and maps will be encoded as JSON obbject
		Map: map[string]string {
			"key": "value",
			"another_key": "another_value",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func template_basic() {
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

func main() {
	context_encoding()
}
