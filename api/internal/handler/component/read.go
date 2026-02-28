package component

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) GetComponentByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing component id")
		return
	}

	var component model.Component
	result := h.DB.First(&component, "id = ?", id)
	if result.Error != nil {
		response.WriteError(w, http.StatusNotFound, "component not found")
		return
	}

	response.WriteJSON(w, http.StatusOK, component)
}

func (h *Handler) GetComponents(w http.ResponseWriter, r *http.Request) {
	var components []model.Component
	result := h.DB.Find(&components)
	if result.Error != nil {
		slog.Error("Error getting components", "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error getting components")
		return
	}

	response.WriteJSON(w, http.StatusOK, components)
}
