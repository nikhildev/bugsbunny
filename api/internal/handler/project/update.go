package project

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
	"github.com/nikhildev/bugsbunny/api/internal/updates"
)

func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "missing project id")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var requestData map[string]any
	if err = json.Unmarshal(body, &requestData); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	u := updates.ExtractUpdates(requestData, model.Project{})
	if len(u) == 0 {
		response.WriteError(w, http.StatusBadRequest, "no fields to update")
		return
	}

	rowsAffected, err := h.Repo.Update(id, u)
	if err != nil {
		slog.Error("Error updating project", "id", id, "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error updating project")
		return
	}
	if rowsAffected == 0 {
		response.WriteError(w, http.StatusNotFound, "project not found")
		return
	}

	updatedProject, err := h.Repo.GetByID(id)
	if err != nil {
		slog.Error("Error fetching updated project", "id", id, "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error fetching updated project")
		return
	}

	if h.VectorStore != nil {
		go func() {
			if err := h.VectorStore.SyncProjectKnowledge(context.Background(), updatedProject.ID, updatedProject.BotKnowledge); err != nil {
				slog.Error("Error syncing project knowledge vectors", "id", updatedProject.ID, "error", err)
			}
		}()
	}

	response.WriteJSON(w, http.StatusOK, updatedProject)
}
