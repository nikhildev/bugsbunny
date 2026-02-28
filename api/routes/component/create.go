package component

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func CreateComponentHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var component models.Component
	if err = json.Unmarshal(body, &component); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		common.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if component.Name == "" || component.Description == "" || component.Owner == "" {
		common.WriteError(w, http.StatusBadRequest, "name, description, and owner are required")
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		slog.Error("Error getting db client", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	result := db.Create(&component)
	if result.Error != nil {
		slog.Error("Error creating component", "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error creating component")
		return
	}

	common.WriteJSON(w, http.StatusCreated, component)
}
