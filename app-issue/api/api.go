package api

import (
	"app-issue/api/handler/health"
	"app-issue/api/handler/issue"
	"app-issue/api/handler/project"
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

	// Project
	projectHandler := project.NewProjectHandler(s.Db)
	router.Handle("GET /projects/", stackNone(api.CreateHandler(projectHandler.GetProjectsPage)))
	router.Handle("GET /projects/{project_id}/", stackNone(api.CreateHandler(projectHandler.GetProjectPage)))
	router.Handle("GET /api/projects/table/", stackNone(api.CreateHandler(projectHandler.GetProjectTable)))
	router.Handle("GET /api/projects/{project_id}/issues/table/", stackNone(api.CreateHandler(projectHandler.GetIssueTable)))

	// Issue
	issueHandler := issue.NewIssueHandler(s.Db)
	router.Handle("GET /issues/{issue_id}/", stackNone(api.CreateHandler(issueHandler.GetIssuePage)))
	router.Handle("GET /api/issues/{issue_id}/comments/table/", stackNone(api.CreateHandler(issueHandler.GetCommentTable)))

	// Health
	healthHandler := health.NewHealthHandler()
	router.Handle("GET /health/live/", stackNone(api.CreateHandler(healthHandler.GetHealthLive)))
	router.Handle("GET /health/ready/", stackNone(api.CreateHandler(healthHandler.GetHealthReady)))
}
