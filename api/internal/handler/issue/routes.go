package issue

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/handler/comment"
	"gorm.io/gorm"
)

func RegisterIssueRoutes(mux *http.ServeMux, db *gorm.DB) {
	h := &Handler{DB: db}
	ch := &comment.Handler{DB: db}
	mux.HandleFunc("GET /issues", h.GetIssues)
	mux.HandleFunc("POST /issues", h.CreateIssue)
	mux.HandleFunc("GET /issues/{id}", h.GetIssueByID)
	mux.HandleFunc("PATCH /issues/{id}", h.UpdateIssue)
	mux.HandleFunc("DELETE /issues/{id}", h.DeleteIssueByID)
	mux.HandleFunc("POST /issues/{id}/comments", ch.CreateComment)
}
