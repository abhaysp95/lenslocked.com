package views

import "text/template"

type View struct {
	Template *template.Template
}

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t := template.Must(template.ParseFiles(files...))

	return &View {
		Template: t,
	}
}
