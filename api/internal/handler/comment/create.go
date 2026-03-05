package comment

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/httputil"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	issueId := r.PathValue("id")
	if issueId == "" {
		httputil.WriteError(w, http.StatusBadRequest, "missing issue id")
		return
	}

	authorID, err := uuid.Parse(r.Header.Get("X-User-UUID"))
	if err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "X-User-UUID header is missing or invalid")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading request body", "error", err)
		httputil.WriteError(w, http.StatusInternalServerError, "error reading request body")
		return
	}

	var c model.Comment
	if err = json.Unmarshal(body, &c); err != nil {
		slog.Error("Error unmarshalling request body", "error", err)
		httputil.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Set fields after unmarshal to prevent them being overwritten by the request body
	c.IssueID = issueId
	c.Author = authorID

	if err = h.Repo.Create(&c); err != nil {
		slog.Error("Error creating comment", "error", err)
		httputil.WriteError(w, http.StatusInternalServerError, "error creating comment")
		return
	}

	httputil.WriteJSON(w, http.StatusCreated, c)
}
