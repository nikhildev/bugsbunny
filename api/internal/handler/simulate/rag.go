package simulate

import (
	"encoding/json"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/response"
	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
)

func (h *Handler) RAG(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.Text == "" {
		response.WriteError(w, http.StatusBadRequest, "text is required")
		return
	}

	vector, err := vectorstore.GetVectorForText(r.Context(), req.Text)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "vectorization failed: "+err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]any{
		"text":       req.Text,
		"vector":     vector,
		"dimensions": len(vector),
	})
}
