package component

import (
	"fmt"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func DeleteComponentHandler(w http.ResponseWriter, r *http.Request) {
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

	result := db.Model(&models.Component{}).Where("id = ?", id).Update("status", models.DELETED)
	if result.Error != nil {
		common.JSONError(w, "internal server error", http.StatusInternalServerError)
		fmt.Println("Error deleting component", result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Component deleted successfully"))
}
