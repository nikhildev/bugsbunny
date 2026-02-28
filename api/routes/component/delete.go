package component

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func (h *Handler) DeleteComponent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing component id")
		return
	}

	// Soft delete: set status to DELETED rather than removing the row
	result := h.DB.Model(&models.Component{}).Where("id = ?", id).Update("status", models.DELETED)
	if result.Error != nil {
		slog.Error("Error deleting component", "id", id, "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error deleting component")
		return
	}

	common.WriteJSON(w, http.StatusOK, map[string]string{"message": "component deleted successfully"})
}
