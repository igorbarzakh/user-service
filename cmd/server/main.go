package main

import (
	"net/http"

	"github.com/igorbarzakh/user-service/internal/platform/http/middleware"
	"github.com/igorbarzakh/user-service/internal/platform/logger"
	"github.com/igorbarzakh/user-service/internal/user"
)

func main() {
	log := logger.New()

	mux := http.NewServeMux()
	
	mux.HandleFunc("/users", user.Handler)

	handler := middleware.Logging(mux)

	log.Println("server started on :8080")

	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		log.Fatal(err)
	}
}