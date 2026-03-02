package project

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing project id")
		return
	}

	var project model.Project
	result := h.DB.First(&project, "id = ?", id)
	if result.Error != nil {
		response.WriteError(w, http.StatusNotFound, "project not found")
		return
	}

	response.WriteJSON(w, http.StatusOK, project)
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	var projects []model.Project
	result := h.DB.Find(&projects)
	if result.Error != nil {
		slog.Error("Error getting projects", "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error getting projects")
		return
	}

	response.WriteJSON(w, http.StatusOK, projects)
}
