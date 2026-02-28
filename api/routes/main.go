package routes

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/routes/comments"
	"github.com/nikhildev/bugsbunny/api/routes/component"
	"github.com/nikhildev/bugsbunny/api/routes/issue"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	RegisterHealthRoutes(mux)
	component.RegisterComponentRoutes(mux)
	issue.RegisterIssueRoutes(mux)
	comments.RegisterCommentsRoutes(mux)
	return mux
}
