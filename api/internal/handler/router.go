package handler

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/handler/issue"
	"github.com/nikhildev/bugsbunny/api/internal/handler/project"
	"github.com/nikhildev/bugsbunny/api/internal/handler/search"
	"github.com/nikhildev/bugsbunny/api/internal/handler/simulate"
	"github.com/nikhildev/bugsbunny/api/internal/repository"
	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, vs *vectorstore.VectorStore) *http.ServeMux {
	projectRepo := repository.NewProjectRepository(db)
	issueRepo := repository.NewIssueRepository(db)
	commentRepo := repository.NewCommentRepository(db)

	mux := http.NewServeMux()
	RegisterHealthRoutes(mux)
	project.RegisterProjectRoutes(mux, projectRepo, vs)
	issue.RegisterIssueRoutes(mux, issueRepo, commentRepo)
	search.RegisterSearchRoutes(mux, vs)
	simulate.RegisterSimulateRoutes(mux, vs)
	return mux
}
