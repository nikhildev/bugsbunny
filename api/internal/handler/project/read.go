package project

import (
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/httputil"
)

func (h *Handler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		httputil.WriteError(w, http.StatusBadRequest, "missing project id")
		return
	}

	project, err := h.Repo.GetByID(id)
	if err != nil {
		httputil.WriteError(w, http.StatusNotFound, "project not found")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, project)
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Repo.GetAll()
	if err != nil {
		slog.Error("Error getting projects", "error", err)
		httputil.WriteError(w, http.StatusInternalServerError, "error getting projects")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, projects)
}
