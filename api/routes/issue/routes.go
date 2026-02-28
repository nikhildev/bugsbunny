package issue

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/routes/comments"
)

func RegisterIssueRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /issues", CreateIssueHandler)
	mux.HandleFunc("GET /issues/{id}", GetIssueByIDHandler)
	mux.HandleFunc("GET /issues", GetIssuesHandler)
	mux.HandleFunc("PATCH /issues/{id}", UpdateIssueHandler)
	mux.HandleFunc("DELETE /issues/{id}", DeleteIssueByIDHandler)
	mux.HandleFunc("POST /issues/{id}/comments", comments.CreateCommentHandler)
}
