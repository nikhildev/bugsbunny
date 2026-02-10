package issue

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/models"
)

func GetIssueByIDHandler(w http.ResponseWriter, r *http.Request) {
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

	var issue models.Issue
	result := db.First(&issue, "id = ?", id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Issue not found", result.Error)
		w.Write([]byte("Issue not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(issue)
}

func GetIssuesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := clients.GetDbClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error getting db client", err)
		return
	}

	var issues []models.Issue
	result := db.Find(&issues)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error getting issues", result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(issues)
}
