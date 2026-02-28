package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func GetIssueByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		slog.Error("Error getting db client", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	var issue models.Issue
	result := db.First(&issue, "id = ?", id)
	if result.Error != nil {
		slog.Error("Issue not found", "error", result.Error)
		common.WriteError(w, http.StatusNotFound, "issue not found")
		return
	}

	common.WriteJSON(w, http.StatusOK, issue)
}

func GetIssuesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := clients.GetDbClient()
	if err != nil {
		slog.Error("Error getting db client", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	var issues []models.Issue
	result := db.Preload("Reporter").Preload("Assignee").Preload("Component").Preload("Collaborators").Preload("CC").Find(&issues)

	if result.Error != nil {
		slog.Error("Error getting issues", "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error getting issues")
		return
	}

	common.WriteJSON(w, http.StatusOK, issues)
}
