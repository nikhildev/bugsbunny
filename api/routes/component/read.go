package component

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func (h *Handler) GetComponentByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.WriteError(w, http.StatusBadRequest, "missing component id")
		return
	}

	var component models.Component
	result := h.DB.First(&component, "id = ?", id)
	if result.Error != nil {
		common.WriteError(w, http.StatusNotFound, "component not found")
		return
	}

	common.WriteJSON(w, http.StatusOK, component)
}

func (h *Handler) GetComponents(w http.ResponseWriter, r *http.Request) {
	var components []models.Component
	result := h.DB.Find(&components)
	if result.Error != nil {
		slog.Error("Error getting components", "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error getting components")
		return
	}

	common.WriteJSON(w, http.StatusOK, components)
}
