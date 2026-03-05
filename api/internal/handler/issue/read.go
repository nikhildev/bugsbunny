package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	issue, err := h.Repo.GetByID(id)
	if err != nil {
		slog.Error("Issue not found", "error", err)
		response.WriteError(w, http.StatusNotFound, "issue not found")
		return
	}

	response.WriteJSON(w, http.StatusOK, issue)
}

func (h *Handler) GetIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := h.Repo.GetAll()
	if err != nil {
		slog.Error("Error getting issues", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error getting issues")
		return
	}

	response.WriteJSON(w, http.StatusOK, issues)
}
