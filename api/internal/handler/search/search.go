package search

import (
	"net/http"
	"strconv"

	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) SearchKnowledge(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		response.WriteError(w, http.StatusBadRequest, "q parameter is required")
		return
	}

	projectID := r.URL.Query().Get("project_id")

	topK := 5
	if tkStr := r.URL.Query().Get("top_k"); tkStr != "" {
		parsed, err := strconv.Atoi(tkStr)
		if err != nil || parsed < 1 {
			response.WriteError(w, http.StatusBadRequest, "top_k must be a positive integer")
			return
		}
		if parsed > 20 {
			parsed = 20
		}
		topK = parsed
	}

	results, err := h.VectorStore.SearchKnowledge(r.Context(), query, topK, projectID)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "search failed: "+err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]any{
		"query":   query,
		"top_k":   topK,
		"results": results,
	})
}
