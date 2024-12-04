package api

import (
	"app-issue/api/handler/fault"
	"app-issue/api/handler/health"
	"app-issue/api/handler/issue"
	"app-issue/api/handler/project"
	"app-issue/api/handler/welcome"
	"database/sql"
	"framework/api"
	"log"
	"net/http"
	"os"
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
	jwtHandler := api.NewJwtHandler(os.Getenv("JWT_NAME"), []byte(os.Getenv("JWT_SECRET")))
	authHandler := api.NewAuthHandler(*jwtHandler, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, os.Getenv("URL_PREFIX_LOGIN")+"/", http.StatusSeeOther)
	})

	// Middleware
	stackLog := api.CreateMiddlewareStack(api.LoggingMiddleware)
	stackLogAuth := api.CreateMiddlewareStack(
		api.LoggingMiddleware,
		authHandler.AuthMiddleware,
	)

	// Welcome
	welcomeHandler := welcome.NewWelcomeHandler()
	router.Handle("GET /{$}", stackLogAuth(api.CreateHandler(welcomeHandler.GetWelcomePage)))

	// Project
	projectHandler := project.NewProjectHandler(s.Db)
	router.Handle("GET /projects/{$}", stackLogAuth(api.CreateHandler(projectHandler.GetProjectsPage)))
	router.Handle("GET /projects/{projectId}/{$}", stackLogAuth(api.CreateHandler(projectHandler.GetProjectPage)))
	router.Handle("GET /api/projects/table/{$}", stackLogAuth(api.CreateHandler(projectHandler.GetProjectsTable)))
	router.Handle("GET /api/projects/{projectId}/issues/table/{$}", stackLogAuth(api.CreateHandler(projectHandler.GetIssuesTable)))
	router.Handle("POST /api/projects/new/{$}", stackLogAuth(api.CreateHandler(projectHandler.PostProjectNew)))

	// Issue
	issueHandler := issue.NewIssueHandler(s.Db)
	router.Handle("GET /issues/{issueId}/{$}", stackLogAuth(api.CreateHandler(issueHandler.GetIssuePage)))
	router.Handle("GET /api/issues/{issueId}/comments/table/{$}", stackLogAuth(api.CreateHandler(issueHandler.GetCommentsTable)))
	router.Handle("POST /api/issues/new/{$}", stackLogAuth(api.CreateHandler(issueHandler.PostIssueNew)))
	router.Handle("POST /api/comments/new/{$}", stackLogAuth(api.CreateHandler(issueHandler.PostCommentNew)))

	// Create new issue form
	router.Handle("GET /api/status/form/{$}", stackLogAuth(api.CreateHandler(issueHandler.GetStatusesForm)))
	router.Handle("GET /api/priority/form/{$}", stackLogAuth(api.CreateHandler(issueHandler.GetPrioritiesForm)))
	router.Handle("GET /api/branch/form/{$}", stackLogAuth(api.CreateHandler(issueHandler.GetBranchesForm)))

	// Health
	healthHandler := health.NewHealthHandler(s.Db)
	router.Handle("GET /health/live/{$}", stackLog(api.CreateHandler(healthHandler.GetHealthLive)))
	router.Handle("GET /health/ready/{$}", stackLog(api.CreateHandler(healthHandler.GetHealthReady)))

	// Fault tolerance testing
	faultTestHandler := fault.NewFaultTestHandler()
	router.Handle("GET /fault/timeout/{$}", stackLog(api.CreateHandler(faultTestHandler.GetTimeoutBad)))
	router.Handle("GET /fault/retry/{$}", stackLog(api.CreateHandler(faultTestHandler.GetRetryBad)))
}
