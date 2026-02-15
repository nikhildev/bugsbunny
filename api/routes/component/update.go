package component

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func UpdateComponentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing component id"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error reading request body", err)
		return
	}

	// Parse request body into a map to detect which fields were provided
	var requestData map[string]any
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error unmarshalling request body", err)
		return
	}

	// Build updates map with only the fields present in the request
	updates := common.ExtractUpdates(requestData, models.Component{})

	// Return error if no fields to update
	if len(updates) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No fields to update"))
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error getting db client", err)
		return
	}

	result := db.Model(&models.Component{}).Where("id = ?", id).Updates(updates)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Component not found"))
		return
	}

	// Fetch the updated component to return it
	var updatedComponent models.Component
	if err := db.Where("id = ?", id).First(&updatedComponent).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error fetching updated component", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedComponent)
}
