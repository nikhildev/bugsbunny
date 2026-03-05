package project

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/repository"
	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
)

func RegisterProjectRoutes(mux *http.ServeMux, repo repository.ProjectRepo, vs *vectorstore.VectorStore) {
	h := &Handler{Repo: repo, VectorStore: vs}
	mux.HandleFunc("GET /projects", h.GetProjects)
	mux.HandleFunc("POST /projects", h.CreateProject)
	mux.HandleFunc("GET /projects/{id}", h.GetProjectByID)
	mux.HandleFunc("DELETE /projects/{id}", h.DeleteProject)
	mux.HandleFunc("PATCH /projects/{id}", h.UpdateProject)
}
