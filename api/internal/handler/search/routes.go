package search

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
)

func RegisterSearchRoutes(mux *http.ServeMux, vs *vectorstore.VectorStore) {
	h := &Handler{VectorStore: vs}
	mux.HandleFunc("GET /search/knowledge", h.SearchKnowledge)
}
