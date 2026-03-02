package project

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing project id")
		return
	}

	result := h.DB.Model(&model.Project{}).Where("id = ?", id).Update("status", model.DELETED)
	if result.Error != nil {
		slog.Error("Error deleting project", "id", id, "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error deleting project")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
