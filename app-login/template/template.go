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
)

func RenderLoginLayout(wr io.Writer, name string, data any) error {
	return template.Must(template.New("layout").Funcs(getFuncMap()).ParseFS(layout, "layout/*.tmpl")).ExecuteTemplate(wr, name, data)
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"urlPrefixLogin": func() template.URL {
			return template.URL(os.Getenv("URL_PREFIX_LOGIN"))
		},
		"urlPrefixStatic": func() template.URL {
			return template.URL(os.Getenv("URL_PREFIX_STATIC"))
		},
	}
}
