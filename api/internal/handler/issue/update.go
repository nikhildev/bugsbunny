package issue

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/httputil"
	"github.com/nikhildev/bugsbunny/api/internal/updates"
)

func (h *Handler) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		httputil.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		httputil.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var requestData map[string]any
	if err = json.Unmarshal(body, &requestData); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	u := updates.ExtractUpdates(requestData, model.Issue{})
	if len(u) == 0 {
		httputil.WriteError(w, http.StatusBadRequest, "no fields to update")
		return
	}

	rowsAffected, err := h.Repo.Update(id, u)
	if err != nil {
		slog.Error("Error updating issue", "id", id, "error", err)
		httputil.WriteError(w, http.StatusInternalServerError, "error updating issue")
		return
	}
	if rowsAffected == 0 {
		httputil.WriteError(w, http.StatusNotFound, "issue not found")
		return
	}

	updatedIssue, err := h.Repo.GetByID(id)
	if err != nil {
		slog.Error("Error fetching updated issue", "id", id, "error", err)
		httputil.WriteError(w, http.StatusInternalServerError, "error fetching updated issue")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, updatedIssue)
}
