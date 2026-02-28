package component

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func CreateComponentHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		common.JSONError(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var component models.Component
	if err = json.Unmarshal(body, &component); err != nil {
		common.JSONError(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	result := db.Create(&component)
	if result.Error != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	common.JSONSuccess(w, http.StatusCreated, component)
}
