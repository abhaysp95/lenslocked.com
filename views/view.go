package views

import (
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	LayoutDir = "views/layouts/"
	TemplateExt = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t := template.Must(template.ParseFiles(files...))

	return &View {
		Template: t,
		Layout: layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
