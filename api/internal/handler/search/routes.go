package search

import "net/http"

func RegisterSearchRoutes(mux *http.ServeMux) {
	h := &Handler{}
	mux.HandleFunc("GET /search/knowledge", h.SearchKnowledge)
}
