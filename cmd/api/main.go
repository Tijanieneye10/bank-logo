package main

import (
	router "bankplay/internal/api/routes"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	router.SetupRouter(mux)

	fmt.Println("Server running on port 8081")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
