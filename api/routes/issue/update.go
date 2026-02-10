package issue

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/models"
)

func UpdateIssueHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing issue id"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error reading request body", err)
		return
	}

	var issue models.Issue
	err = json.Unmarshal(body, &issue)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error unmarshalling request body", err)
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error getting db client", err)
		return
	}

	result := db.Model(&models.Issue{}).Where("id = ?", id).Updates(issue)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error updating issue", result.Error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(issue)
}
