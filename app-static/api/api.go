package api

import (
	"app-static/api/handler/health"
	"embed"
	"framework/api"
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

	s.registerHandlers(router)

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

func (s *ApiServer) registerHandlers(router *http.ServeMux) {
	// Middleware
	stackNone := api.CreateMiddlewareStack(api.LoggingMiddleware)

	// Health
	healthHandler := health.NewHealthHandler()
	router.Handle("/health/live", stackNone(api.CreateHandler(healthHandler.GetHealthLive)))
	router.Handle("/health/ready", stackNone(api.CreateHandler(healthHandler.GetHealthReady)))
}
