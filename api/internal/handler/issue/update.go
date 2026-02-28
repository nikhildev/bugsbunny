package issue

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
	"github.com/nikhildev/bugsbunny/api/internal/updates"
)

func (h *Handler) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing issue id")
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

	u := updates.ExtractUpdates(requestData, model.Issue{})
	if len(u) == 0 {
		response.WriteError(w, http.StatusBadRequest, "no fields to update")
		return
	}

	result := h.DB.Model(&model.Issue{}).Where("id = ?", id).Updates(u)
	if result.Error != nil {
		slog.Error("Error updating issue", "id", id, "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error updating issue")
		return
	}
	if result.RowsAffected == 0 {
		response.WriteError(w, http.StatusNotFound, "issue not found")
		return
	}

	var updatedIssue model.Issue
	if err = h.DB.Where("id = ?", id).First(&updatedIssue).Error; err != nil {
		slog.Error("Error fetching updated issue", "id", id, "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error fetching updated issue")
		return
	}

	response.WriteJSON(w, http.StatusOK, updatedIssue)
}
