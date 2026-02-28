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

	// Soft delete: set status to ISSUE_DELETED rather than removing the row
	result := h.DB.Model(&models.Issue{}).Where("id = ?", id).Update("status", models.ISSUE_DELETED)
	if result.Error != nil {
		slog.Error("Error deleting issue", "id", id, "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error deleting issue")
		return
	}

	common.WriteJSON(w, http.StatusOK, map[string]string{"message": "issue deleted successfully"})
}
