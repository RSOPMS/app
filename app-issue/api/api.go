package api

import (
	"app-issue/api/handler/health"
	"app-issue/api/handler/issue"
	"database/sql"
	"framework/api"
	"log"
	"net/http"
)

type ApiServer struct {
	Addr string
	Db   *sql.DB
}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

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

	// Health
	healthHandler := health.NewHealthHandler()
	router.Handle("GET /health/live", stackNone(api.CreateHandler(healthHandler.GetHealthLive)))
	router.Handle("GET /health/ready", stackNone(api.CreateHandler(healthHandler.GetHealthReady)))
}
