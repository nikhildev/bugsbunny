package project

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
	"gorm.io/gorm"
)

func RegisterProjectRoutes(mux *http.ServeMux, db *gorm.DB) {
	_, weaviateErr := vectorstore.GetWeaviateClient()
	h := &Handler{DB: db, VectorSyncEnabled: weaviateErr == nil}
	mux.HandleFunc("GET /projects", h.GetProjects)
	mux.HandleFunc("POST /projects", h.CreateProject)
	mux.HandleFunc("GET /projects/{id}", h.GetProjectByID)
	mux.HandleFunc("DELETE /projects/{id}", h.DeleteProject)
	mux.HandleFunc("PATCH /projects/{id}", h.UpdateProject)
}
