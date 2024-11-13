package api

import (
	"app-issue/api/handler/issue"
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"framework/api"
)

type ApiServer struct {
	Addr string
	Db   *sql.DB
}

//go:embed static
var static embed.FS

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	staticDir, err := fs.Sub(static, "static")
	if err != nil {
		return err
	}

	router.Handle("/", http.FileServer(http.FS(staticDir)))
	s.registerHandlers(router)

	server := http.Server{
		Addr:    s.Addr,
		Handler: router,
	}

	log.Println("Server is listening on", server.Addr)

	return server.ListenAndServe()
}

func (s *ApiServer) registerHandlers(router *http.ServeMux) {
	// Middleware
	stackNone := api.CreateMiddlewareStack(api.LoggingMiddleware)

	// Issue
	issueHandler := issue.NewIssueHandler(s.Db)
	router.Handle("GET /issue", stackNone(api.CreateHandler(issueHandler.GetIssue)))
	router.Handle("GET /api/issue-list", stackNone(api.CreateHandler(issueHandler.GetIssueList)))
}
