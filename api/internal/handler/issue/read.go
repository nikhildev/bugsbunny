package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	var issue model.Issue
	result := h.DB.First(&issue, "id = ?", id)
	if result.Error != nil {
		slog.Error("Issue not found", "error", result.Error)
		response.WriteError(w, http.StatusNotFound, "issue not found")
		return
	}

	response.WriteJSON(w, http.StatusOK, issue)
}

func (h *Handler) GetIssues(w http.ResponseWriter, r *http.Request) {
	var issues []model.Issue
	result := h.DB.Preload("Reporter").Preload("Assignee").Preload("Component").Preload("Collaborators").Preload("CC").Find(&issues)
	if result.Error != nil {
		slog.Error("Error getting issues", "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error getting issues")
		return
	}

	response.WriteJSON(w, http.StatusOK, issues)
}
