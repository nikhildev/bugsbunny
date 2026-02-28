package issue

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/routes/comments"
	"gorm.io/gorm"
)

func RegisterIssueRoutes(mux *http.ServeMux, db *gorm.DB) {
	h := &Handler{DB: db}
	ch := &comments.Handler{DB: db}
	mux.HandleFunc("GET /issues", h.GetIssues)
	mux.HandleFunc("POST /issues", h.CreateIssue)
	mux.HandleFunc("GET /issues/{id}", h.GetIssueByID)
	mux.HandleFunc("PUT /issues/{id}", h.UpdateIssue)
	mux.HandleFunc("DELETE /issues/{id}", h.DeleteIssueByID)
	mux.HandleFunc("POST /issues/{id}/comments", ch.CreateComment)
}
