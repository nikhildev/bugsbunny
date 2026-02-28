package comments

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/models"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {

	issueId := r.PathValue("id")
	if issueId == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Missing issue id")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error reading request body", err)
		return
	}

	var comment models.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Invalid request body", err)
		return
	}

	comment.IssueID = issueId
	comment.Author, err = uuid.Parse(r.Header.Get("X-User-UUID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Invalid user UUID", err)
		return
	}

	db, err := clients.GetDbClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error getting db client", err)
		return
	}

	result := db.Create(&comment)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error creating comment", result.Error.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Comment created successfully", result.RowsAffected)
	json.NewEncoder(w).Encode(result.RowsAffected)

}
