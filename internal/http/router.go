package http

import (
	"net/http"

	"github.com/igorbarzakh/user-service/internal/health"
	"github.com/igorbarzakh/user-service/internal/http/middleware"
	"github.com/igorbarzakh/user-service/internal/platform/httpx"
	"github.com/igorbarzakh/user-service/internal/user"
)

func NewRouter() http.Handler {

	mux := http.NewServeMux()

	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	mux.HandleFunc("GET /health", health.Handler)
	mux.HandleFunc("GET /users", handler.ListUsers)
	mux.HandleFunc("GET /users/{id}", handler.GetUserByID)
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("DELETE /users/{id}", handler.DeleteUser)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {                                                      
		httpx.WriteError(w, "not found", http.StatusNotFound)
})

	return middleware.Logging(mux)
}