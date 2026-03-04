package http

import (
	"net/http"

	"github.com/igorbarzakh/user-service/internal/health"
	"github.com/igorbarzakh/user-service/internal/http/middleware"
	"github.com/igorbarzakh/user-service/internal/user"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/health", health.Handler)
	mux.HandleFunc("/users", user.Handler)

	handler := middleware.Logging(mux)

	return handler
}