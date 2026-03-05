package project

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
)

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var project model.Project
	if err = json.Unmarshal(body, &project); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if project.Name == "" || project.Description == "" || project.Owner == "" {
		response.WriteError(w, http.StatusBadRequest, "name, description, and owner are required")
		return
	}

	if err = h.Repo.Create(&project); err != nil {
		slog.Error("Error creating project", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error creating project")
		return
	}

	if h.VectorStore != nil && len(project.BotKnowledge) > 0 {
		go func() {
			if err := h.VectorStore.SyncProjectKnowledge(context.Background(), project.ID, project.BotKnowledge); err != nil {
				slog.Error("Error syncing project knowledge vectors", "id", project.ID, "error", err)
			}
		}()
	}

	response.WriteJSON(w, http.StatusCreated, project)
}
