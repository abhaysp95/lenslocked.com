package views

import "text/template"

type View struct {
	Template *template.Template
	Layout string
}

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/bootstrap.gohtml",
		"views/layouts/footer.gohtml",
		"views/layouts/navbar.gohtml",
	)
	t := template.Must(template.ParseFiles(files...))

	return &View {
		Template: t,
		Layout: layout,
	}
}
