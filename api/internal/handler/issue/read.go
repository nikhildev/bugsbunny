package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/httputil"
)

func (h *Handler) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		httputil.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	issue, err := h.Repo.GetByID(id)
	if err != nil {
		slog.Error("Issue not found", "error", err)
		httputil.WriteError(w, http.StatusNotFound, "issue not found")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, issue)
}

func (h *Handler) GetIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := h.Repo.GetAll()
	if err != nil {
		slog.Error("Error getting issues", "error", err)
		httputil.WriteError(w, http.StatusInternalServerError, "error getting issues")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, issues)
}
