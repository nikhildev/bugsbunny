package routes

import (
	"net/http"

	// "github.com/nikhildev/bugsbunny/api/routes/component"
	"github.com/nikhildev/bugsbunny/api/routes/issue"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	// component.RegisterRoutes(mux)
	issue.RegisterIssueRoutes(mux)
	return mux
}
