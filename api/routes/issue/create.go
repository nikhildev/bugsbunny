package issue

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func (h *Handler) CreateIssue(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var issue models.Issue
	if err = json.Unmarshal(body, &issue); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		common.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if issue.Title == "" || issue.Description == "" || issue.ReporterId == "" || issue.ComponentID == "" {
		common.WriteError(w, http.StatusBadRequest, "title, description, reporter_id, and component_id are required")
		return
	}

	result := h.DB.Create(&issue)
	if result.Error != nil {
		slog.Error("Error creating issue", "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error creating issue")
		return
	}

	common.WriteJSON(w, http.StatusCreated, issue)
}
