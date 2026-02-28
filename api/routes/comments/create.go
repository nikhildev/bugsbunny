package comments

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	issueId := r.PathValue("id")
	if issueId == "" {
		common.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	authorID, err := uuid.Parse(r.Header.Get("X-User-UUID"))
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, "X-User-UUID header is missing or invalid")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		common.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var comment models.Comment
	if err = json.Unmarshal(body, &comment); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		common.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if comment.Content == "" {
		common.WriteError(w, http.StatusBadRequest, "content is required")
		return
	}

	// Set fields after unmarshal to prevent them being overwritten by the request body
	comment.IssueID = issueId
	comment.Author = authorID

	result := h.DB.Create(&comment)
	if result.Error != nil {
		slog.Error("Error creating comment", "error", result.Error)
		common.WriteError(w, http.StatusInternalServerError, "error creating comment")
		return
	}

	slog.Info("Comment created", "issue_id", issueId, "rows_affected", result.RowsAffected)
	common.WriteJSON(w, http.StatusCreated, comment)
}
