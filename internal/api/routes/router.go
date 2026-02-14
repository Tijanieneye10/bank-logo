package routes

import (
	"bankplay/internal/api/handlers"
	"net/http"
)

func SetupRouter(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.GetBanks)
	mux.HandleFunc("/banks/{code}", handlers.GetSingleBank)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK, its working "))
	})
}
