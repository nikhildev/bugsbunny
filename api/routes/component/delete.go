package component

import (
	"fmt"
	"net/http"

	"github.com/nikhildev/bugsbunny/clients"
	"github.com/nikhildev/bugsbunny/models"
)

func DeleteComponentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing component id"))
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error getting db client", err)
		return
	}

	result := db.Model(&models.Component{}).Where("id = ?", id).Update("status", models.DELETED)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error deleting component", result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Component deleted successfully"))
}
