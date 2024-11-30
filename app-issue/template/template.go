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

	//go:embed project
	project embed.FS

	//go:embed comments
	comments embed.FS
)

func RenderLayout(wr io.Writer, name string, data any) error {
	return template.Must(template.New("layout").ParseFS(layout, "layout/*.tmpl")).ExecuteTemplate(wr, name, data)
}

func RenderIssue(wr io.Writer, name string, data any) error {
	return template.Must(template.New("issue").ParseFS(issue, "issue/*.tmpl")).ExecuteTemplate(wr, name, data)
}

func RenderProject(wr io.Writer, name string, data any) error {
	return template.Must(template.New("project").ParseFS(project, "project/*.tmpl")).ExecuteTemplate(wr, name, data)
}

func RenderComments(wr io.Writer, name string, data any) error {
	return template.Must(template.New("comments").ParseFS(comments, "comments/*.tmpl")).ExecuteTemplate(wr, name, data)
}
