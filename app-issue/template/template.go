package template

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed layout
	layout embed.FS

	//go:embed issue
	issue embed.FS
)

func RenderLayout(wr io.Writer, name string, data any) error {
	return template.Must(template.New("layout").ParseFS(layout, "layout/*.tmpl")).ExecuteTemplate(wr, name, data)
}

func RenderIssue(wr io.Writer, name string, data any) error {
	return template.Must(template.New("issue").ParseFS(issue, "issue/*.tmpl")).ExecuteTemplate(wr, name, data)
}
