package simulate

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
)

func RegisterSimulateRoutes(mux *http.ServeMux, vs *vectorstore.VectorStore) {
	h := &Handler{VectorStore: vs}
	mux.HandleFunc("POST /simulate/rag", h.RAG)
}
