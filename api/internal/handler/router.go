package handler

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/handler/issue"
	"github.com/nikhildev/bugsbunny/api/internal/handler/project"
	"github.com/nikhildev/bugsbunny/api/internal/handler/search"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()
	RegisterHealthRoutes(mux)
	project.RegisterProjectRoutes(mux, db)
	issue.RegisterIssueRoutes(mux, db)
	search.RegisterSearchRoutes(mux)
	return mux
}
