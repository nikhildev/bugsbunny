package issue

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) CreateIssue(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var issue model.Issue
	if err = json.Unmarshal(body, &issue); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if issue.Title == "" || issue.Description == "" || issue.ReporterId == "" || issue.ComponentID == "" {
		response.WriteError(w, http.StatusBadRequest, "title, description, reporter_id, and component_id are required")
		return
	}

	result := h.DB.Create(&issue)
	if result.Error != nil {
		slog.Error("Error creating issue", "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error creating issue")
		return
	}

	response.WriteJSON(w, http.StatusCreated, issue)
}
