package routes

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})
	return mux
}
