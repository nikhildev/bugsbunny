package issue

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/handler/comment"
	"github.com/nikhildev/bugsbunny/api/internal/repository"
)

func RegisterIssueRoutes(mux *http.ServeMux, issueRepo repository.IssueRepo, commentRepo repository.CommentRepo) {
	h := &Handler{Repo: issueRepo}
	ch := &comment.Handler{Repo: commentRepo}
	mux.HandleFunc("GET /issues", h.GetIssues)
	mux.HandleFunc("POST /issues", h.CreateIssue)
	mux.HandleFunc("GET /issues/{id}", h.GetIssueByID)
	mux.HandleFunc("PATCH /issues/{id}", h.UpdateIssue)
	mux.HandleFunc("DELETE /issues/{id}", h.DeleteIssueByID)
	mux.HandleFunc("POST /issues/{id}/comments", ch.CreateComment)
}
