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
		common.JSONError(w, "Missing component id", http.StatusBadRequest)
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		slog.Error("Error getting db client", "error", err)
		return
	}

	var component models.Component
	result := db.First(&component, "id = ?", id)
	if result.Error != nil {
		common.JSONError(w, "Component not found", http.StatusNotFound)
		slog.Error("Component not found", "error", result.Error)
		return
	}

	common.JSONSuccess(w, http.StatusOK, component)
}

func GetComponentsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := clients.GetDbClient()
	if err != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		slog.Error("Error getting db client", "error", err)
		return
	}

	var components []models.Component
	result := db.Find(&components)
	if result.Error != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		slog.Error("Error getting components", "error", result.Error)
		return
	}

	common.JSONSuccess(w, http.StatusOK, components)
}
