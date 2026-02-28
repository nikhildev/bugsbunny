package routes

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/routes/component"
	"github.com/nikhildev/bugsbunny/api/routes/issue"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()
	RegisterHealthRoutes(mux)
	component.RegisterComponentRoutes(mux, db)
	issue.RegisterIssueRoutes(mux, db)
	return mux
}
