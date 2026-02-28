package issue

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func DeleteIssueByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		slog.Error("Error getting db client", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	// Soft delete: set status to DELETED rather than removing the row
	result := db.Model(&models.Issue{}).Where("id = ?", id).Update("status", models.DELETED)
	if result.Error != nil {
		slog.Error("Error deleting issue", "id", id, "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error deleting issue")
		return
	}

	common.WriteJSON(w, http.StatusOK, map[string]string{"message": "issue deleted successfully"})
}
