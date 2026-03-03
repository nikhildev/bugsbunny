package simulate

import "net/http"

func RegisterSimulateRoutes(mux *http.ServeMux) {
	h := &Handler{}
	mux.HandleFunc("POST /simulate/rag", h.RAG)
}
