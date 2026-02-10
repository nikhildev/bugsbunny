package routes

import "net/http"

func RegisterHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", HealthHandler)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
