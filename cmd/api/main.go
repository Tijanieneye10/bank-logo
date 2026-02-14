package main

import (
	router "bankplay/internal/api"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	router.SetupRouter(mux)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		return
	}
}
