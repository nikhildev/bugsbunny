package project

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing project id")
		return
	}

	project, err := h.Repo.GetByID(id)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "project not found")
		return
	}

	response.WriteJSON(w, http.StatusOK, project)
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Repo.GetAll()
	if err != nil {
		slog.Error("Error getting projects", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error getting projects")
		return
	}

	response.WriteJSON(w, http.StatusOK, projects)
}
