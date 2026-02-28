package component

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func UpdateComponentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing component id")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var requestData map[string]any
	if err = json.Unmarshal(body, &requestData); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		common.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	updates := common.ExtractUpdates(requestData, models.Component{})
	if len(updates) == 0 {
		common.WriteError(w, http.StatusBadRequest, "no fields to update")
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		slog.Error("Error getting db client", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	result := db.Model(&models.Component{}).Where("id = ?", id).Updates(updates)
	if result.RowsAffected == 0 {
		common.WriteError(w, http.StatusNotFound, "component not found")
		return
	}

	var updatedComponent models.Component
	if err = db.Where("id = ?", id).First(&updatedComponent).Error; err != nil {
		slog.Error("Error fetching updated component", "id", id, "error", err)
		common.WriteError(w, http.StatusInternalServerError, "error fetching updated component")
		return
	}

	common.WriteJSON(w, http.StatusOK, updatedComponent)
}
