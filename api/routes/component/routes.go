package component

import "net/http"

func RegisterComponentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /components", GetComponentsHandler)
	mux.HandleFunc("POST /components", CreateComponentHandler)
	mux.HandleFunc("GET /components/{id}", GetComponentByIDHandler)
	mux.HandleFunc("DELETE /components/{id}", DeleteComponentHandler)
	mux.HandleFunc("PATCH /components/{id}", UpdateComponentHandler)
}
