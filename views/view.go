package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	ViewDir     = "views/"
	LayoutDir   = "views/layouts/"
	TemplateExt = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func NewNamedView(layout string, view string) *View {
	var viewFile = ViewDir + view + TemplateExt
	namedView := NewView(layout, viewFile)
	return namedView
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
