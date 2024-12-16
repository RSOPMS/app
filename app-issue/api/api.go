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
	stackLog := api.CreateMiddlewareStack(api.LoggingMiddleware)

	// Project
	projectHandler := project.NewProjectHandler(s.Db)
	router.Handle("GET /projects", stackLog(api.CreateHandler(projectHandler.GetProjectsPage)))
	router.Handle("GET /projects/{projectId}", stackLog(api.CreateHandler(projectHandler.GetProjectPage)))
	router.Handle("GET /api/projects/table", stackLog(api.CreateHandler(projectHandler.GetProjectsTable)))
	router.Handle("GET /api/projects/{projectId}/issues/table", stackLog(api.CreateHandler(projectHandler.GetIssuesTable)))

	// Issue
	issueHandler := issue.NewIssueHandler(s.Db)
	router.Handle("GET /issues/{issueId}", stackLog(api.CreateHandler(issueHandler.GetIssuePage)))
	router.Handle("GET /api/issues/{issueId}/comments/table", stackLog(api.CreateHandler(issueHandler.GetCommentsTable)))
	router.Handle("POST /api/issues/{$}", stackLog(api.CreateHandler(issueHandler.CreateNewIssue)))

	// Create new issue form
	router.Handle("GET /api/status-form/{$}", stackLog(api.CreateHandler(issueHandler.GetStatusesForm)))
	router.Handle("GET /api/priority-form/{$}", stackLog(api.CreateHandler(issueHandler.GetPrioritiesForm)))
	router.Handle("GET /api/branch-form/{$}", stackLog(api.CreateHandler(issueHandler.GetBranchesForm)))

	// Health
	healthHandler := health.NewHealthHandler(s.Db)
	router.Handle("GET /health/live", stackLog(api.CreateHandler(healthHandler.GetHealthLive)))
	router.Handle("GET /health/ready", stackLog(api.CreateHandler(healthHandler.GetHealthReady)))
}
