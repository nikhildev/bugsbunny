package issue

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func UpdateIssueHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var requestData map[string]any
	if err = json.Unmarshal(body, &requestData); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		common.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	updates := common.ExtractUpdates(requestData, models.Issue{})
	if len(updates) == 0 {
		common.WriteError(w, http.StatusBadRequest, "no fields to update")
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		slog.Error("Error getting db client", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	result := db.Model(&models.Issue{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		slog.Error("Error updating issue", "id", id, "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error updating issue")
		return
	}
	if result.RowsAffected == 0 {
		common.WriteError(w, http.StatusNotFound, "issue not found")
		return
	}

	var updatedIssue models.Issue
	if err = db.Where("id = ?", id).First(&updatedIssue).Error; err != nil {
		slog.Error("Error fetching updated issue", "id", id, "error", err)
		common.WriteError(w, http.StatusInternalServerError, "error fetching updated issue")
		return
	}

	common.WriteJSON(w, http.StatusOK, updatedIssue)
}
