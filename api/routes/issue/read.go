package issue

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func GetIssueByIDHandler(w http.ResponseWriter, r *http.Request) {
	common.EnableCors(w)
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing issue id"))
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting db client", "error", err)
		return
	}

	var issue models.Issue
	result := db.First(&issue, "id = ?", id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		slog.Error("Issue not found", "error", result.Error)
		w.Write([]byte("Issue not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(issue); err != nil {
		slog.Error("Failed to encode issue", "error", err)
	}
}

func GetIssuesHandler(w http.ResponseWriter, r *http.Request) {
	common.EnableCors(w)
	db, err := clients.GetDbClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting db client", "error", err)
		return
	}

	var issues []models.Issue
	result := db.Preload("Reporter").Preload("Assignee").Preload("Component").Preload("Collaborators").Preload("CC").Find(&issues)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting issues", "error", result.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(issues); err != nil {
		slog.Error("Failed to encode issues", "error", err)
	}
}
