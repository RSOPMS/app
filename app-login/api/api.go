package api

import (
	"app-login/api/handler/health"
	"app-login/api/handler/login"
	"app-login/api/handler/register"
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

	// Login
	loginHandler := login.NewLoginHandler(s.Db)
	router.Handle("GET /{$}", stackLog(api.CreateHandler(loginHandler.GetLoginPage)))
	router.Handle("POST /{$}", stackLog(api.CreateHandler(loginHandler.ProcessLogin)))
	router.Handle("GET /logout/{$}", stackLog(api.CreateHandler(loginHandler.ProcessLogout)))

	// Register
	registerHandler := register.NewRegisterHandler(s.Db)
	router.Handle("GET /register/{$}", stackLog(api.CreateHandler(registerHandler.GetRegisterPage)))
	router.Handle("POST /register/{$}", stackLog(api.CreateHandler(registerHandler.PostRegisterNew)))
	router.Handle("GET /api/user/new/{$}", stackLog(api.CreateHandler(registerHandler.GetRolesForm)))

	// Health
	healthHandler := health.NewHealthHandler(s.Db)
	router.Handle("GET /health/live/{$}", stackLog(api.CreateHandler(healthHandler.GetHealthLive)))
	router.Handle("GET /health/ready/{$}", stackLog(api.CreateHandler(healthHandler.GetHealthReady)))
}
