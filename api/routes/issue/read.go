package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func (h *Handler) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	common.EnableCors(w)
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	var issue models.Issue
	result := h.DB.First(&issue, "id = ?", id)
	if result.Error != nil {
		slog.Error("Issue not found", "error", result.Error)
		common.WriteError(w, http.StatusNotFound, "issue not found")
		return
	}

	common.WriteJSON(w, http.StatusOK, issue)
}

func (h *Handler) GetIssues(w http.ResponseWriter, r *http.Request) {
	common.EnableCors(w)

	var issues []models.Issue
	result := h.DB.Preload("Reporter").Preload("Assignee").Preload("Component").Preload("Collaborators").Preload("CC").Find(&issues)

	if result.Error != nil {
		slog.Error("Error getting issues", "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error getting issues")
		return
	}

	common.WriteJSON(w, http.StatusOK, issues)
}
