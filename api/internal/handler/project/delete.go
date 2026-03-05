package project

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing project id")
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		slog.Error("Error deleting project", "id", id, "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error deleting project")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
