package component

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) DeleteComponent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing component id")
		return
	}

	result := h.DB.Model(&model.Component{}).Where("id = ?", id).Update("status", model.DELETED)
	if result.Error != nil {
		slog.Error("Error deleting component", "id", id, "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error deleting component")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
