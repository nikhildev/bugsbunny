package component

import (
	"fmt"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func GetComponentByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		common.JSONError(w, "Missing component id", http.StatusBadRequest)
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		fmt.Println("Error getting db client", err)
		return
	}

	var component models.Component
	result := db.First(&component, "id = ?", id)
	if result.Error != nil {
		common.JSONError(w, "Component not found", http.StatusNotFound)
		fmt.Println("Component not found", result.Error)
		return
	}

	common.JSONSuccess(w, http.StatusOK, component)
}

func GetComponentsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := clients.GetDbClient()
	if err != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		fmt.Println("Error getting db client", err)
		return
	}

	var components []models.Component
	result := db.Find(&components)
	if result.Error != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		fmt.Println("Error getting components", result.Error)
		return
	}

	common.JSONSuccess(w, http.StatusOK, components)
}
