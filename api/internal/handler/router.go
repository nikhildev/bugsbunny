package handler

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/handler/issue"
	"github.com/nikhildev/bugsbunny/api/internal/handler/project"
	"github.com/nikhildev/bugsbunny/api/internal/handler/search"
	"github.com/nikhildev/bugsbunny/api/internal/handler/simulate"
	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, vs *vectorstore.VectorStore) *http.ServeMux {
	mux := http.NewServeMux()
	RegisterHealthRoutes(mux)
	project.RegisterProjectRoutes(mux, db, vs)
	issue.RegisterIssueRoutes(mux, db)
	search.RegisterSearchRoutes(mux, vs)
	simulate.RegisterSimulateRoutes(mux, vs)
	return mux
}
