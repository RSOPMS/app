package api

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

type ApiServer struct {
	Addr string
}

//go:embed static
var static embed.FS

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	staticDir, err := fs.Sub(static, "static")
	if err != nil {
		return err
	}

	router.Handle("/", http.FileServerFS(staticDir))

	server := http.Server{
		Addr:    s.Addr,
		Handler: router,
	}

	log.Println("Server is listening on", server.Addr)

	return server.ListenAndServe()
}
