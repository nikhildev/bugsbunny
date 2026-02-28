package component

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func GetComponentByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing component id")
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		slog.Error("Error getting db client", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	var component models.Component
	result := db.First(&component, "id = ?", id)
	if result.Error != nil {
		common.WriteError(w, http.StatusNotFound, "component not found")
		return
	}

	common.WriteJSON(w, http.StatusOK, component)
}

func GetComponentsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := clients.GetDbClient()
	if err != nil {
		slog.Error("Error getting db client", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	var components []models.Component
	result := db.Find(&components)
	if result.Error != nil {
		slog.Error("Error getting components", "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error getting components")
		return
	}

	common.WriteJSON(w, http.StatusOK, components)
}
