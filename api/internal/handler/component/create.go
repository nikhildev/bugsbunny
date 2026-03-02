package component

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/response"
	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
)

func (h *Handler) CreateComponent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var component model.Component
	if err = json.Unmarshal(body, &component); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if component.Name == "" || component.Description == "" || component.Owner == "" {
		response.WriteError(w, http.StatusBadRequest, "name, description, and owner are required")
		return
	}

	result := h.DB.Create(&component)
	if result.Error != nil {
		slog.Error("Error creating component", "error", result.Error)
		response.WriteError(w, http.StatusInternalServerError, "error creating component")
		return
	}

	if h.VectorSyncEnabled && len(component.BotKnowledge) > 0 {
		go func() {
			if err := vectorstore.SyncComponentKnowledge(context.Background(), component.ID, component.BotKnowledge); err != nil {
				slog.Error("Error syncing component knowledge vectors", "id", component.ID, "error", err)
			}
		}()
	}

	response.WriteJSON(w, http.StatusCreated, component)
}
