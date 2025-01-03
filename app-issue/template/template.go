package template

import (
	"embed"
	"html/template"
	"io"
	"os"
)

var (
	//go:embed layout
	layout embed.FS

	//go:embed issue
	issue embed.FS

	//go:embed project
	project embed.FS
)

func RenderLayout(wr io.Writer, name string, data any) error {
	return template.Must(template.New("layout").Funcs(getFuncMap()).ParseFS(layout, "layout/*.tmpl")).ExecuteTemplate(wr, name, data)
}

func RenderIssue(wr io.Writer, name string, data any) error {
	return template.Must(template.New("issue").Funcs(getFuncMap()).ParseFS(issue, "issue/*.tmpl")).ExecuteTemplate(wr, name, data)
}

func RenderProject(wr io.Writer, name string, data any) error {
	return template.Must(template.New("project").Funcs(getFuncMap()).ParseFS(project, "project/*.tmpl")).ExecuteTemplate(wr, name, data)
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"urlPrefixIssue": func() template.URL {
			return template.URL(os.Getenv("URL_PREFIX_ISSUE"))
		},
		"urlPrefixLogin": func() template.URL {
			return template.URL(os.Getenv("URL_PREFIX_LOGIN"))
		},
		"urlPrefixStatic": func() template.URL {
			return template.URL(os.Getenv("URL_PREFIX_STATIC"))
		},
	}
}
