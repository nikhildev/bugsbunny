package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) DeleteIssueByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	result := h.DB.Model(&model.Issue{}).Where("id = ?", id).Update("status", model.ISSUE_DELETED)
	if result.Error != nil {
		slog.Error("Error deleting issue", "id", id, "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error deleting issue")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
