package routes

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/routes/comments"
	"github.com/nikhildev/bugsbunny/api/routes/component"
	"github.com/nikhildev/bugsbunny/api/routes/issue"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()
	component.RegisterComponentRoutes(mux, db)
	issue.RegisterIssueRoutes(mux, db)
	comments.RegisterCommentsRoutes(mux, db)
	RegisterHealthRoutes(mux)
	return mux
}
