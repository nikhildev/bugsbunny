package issue

import (
	"net/http"
)

func RegisterIssueRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /issues", CreateIssueHandler)
	mux.HandleFunc("GET /issues/{id}", GetIssueByIDHandler)
	mux.HandleFunc("GET /issues", GetIssuesHandler)
	mux.HandleFunc("PUT /issues/{id}", UpdateIssueHandler)
	mux.HandleFunc("DELETE /issues/{id}", DeleteIssueByIDHandler)
}
