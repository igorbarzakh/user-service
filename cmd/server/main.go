package main

import (
	"net/http"

	apphttp "github.com/igorbarzakh/user-service/internal/http"
	"github.com/igorbarzakh/user-service/internal/platform/logger"
)

func main() {
	log := logger.New()

	handler := apphttp.NewRouter()

	log.Println("server started on :8080")

	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		log.Fatal(err)
	}
}