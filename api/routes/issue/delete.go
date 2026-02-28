package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func (h *Handler) DeleteIssueByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	result := h.DB.Model(&models.Issue{}).Where("id = ?", id).Update("status", models.ISSUE_DELETED)
	if result.Error != nil {
		slog.Error("Error deleting issue", "id", id, "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error deleting issue")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
