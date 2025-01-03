package api

import (
	"app-bulk/api/handler/bulk"
	"app-bulk/api/handler/health"
	"framework/api"
	"log"
	"net/http"
)

type ApiServer struct {
	Addr string
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

	// Bulk
	bulkHandler := bulk.NewBulkHandler()
	router.Handle("POST /api/bulk/{$}", stackLog(api.CreateHandler(bulkHandler.PostBulk)))

	// Health
	healthHandler := health.NewHealthHandler()
	router.Handle("GET /health/live/{$}", stackLog(api.CreateHandler(healthHandler.GetHealthLive)))
	router.Handle("GET /health/ready/{$}", stackLog(api.CreateHandler(healthHandler.GetHealthReady)))
}
