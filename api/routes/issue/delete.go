package issue

import (
	"fmt"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/models"
)

func DeleteIssueByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing issue id"))
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error getting db client", err)
		return
	}

	result := db.Model(&models.Issue{}).Where("id = ?", id).Update("status", models.DELETED)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error deleting issue", result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Issue deleted successfully"))
}
