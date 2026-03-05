package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) DeleteIssueByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		slog.Error("Error deleting issue", "id", id, "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error deleting issue")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
