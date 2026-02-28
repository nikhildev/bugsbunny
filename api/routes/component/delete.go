package component

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func DeleteComponentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.JSONError(w, "Missing component id", http.StatusBadRequest)
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		slog.Error("Error getting db client", "error", err)
		return
	}

	result := db.Model(&models.Component{}).Where("id = ?", id).Update("status", models.DELETED)
	if result.Error != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		slog.Error("Error deleting component", "error", result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Component deleted successfully"))
}
