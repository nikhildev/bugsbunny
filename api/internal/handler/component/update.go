package component

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
	"github.com/nikhildev/bugsbunny/api/internal/updates"
)

func (h *Handler) UpdateComponent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing component id")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var requestData map[string]any
	if err = json.Unmarshal(body, &requestData); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	u := updates.ExtractUpdates(requestData, model.Component{})
	if len(u) == 0 {
		response.WriteError(w, http.StatusBadRequest, "no fields to update")
		return
	}

	result := h.DB.Model(&model.Component{}).Where("id = ?", id).Updates(u)
	if result.RowsAffected == 0 {
		response.WriteError(w, http.StatusNotFound, "component not found")
		return
	}

	var updatedComponent model.Component
	if err = h.DB.Where("id = ?", id).First(&updatedComponent).Error; err != nil {
		slog.Error("Error fetching updated component", "id", id, "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error fetching updated component")
		return
	}

	response.WriteJSON(w, http.StatusOK, updatedComponent)
}
