package component

import (
	"net/http"

	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
	"gorm.io/gorm"
)

func RegisterComponentRoutes(mux *http.ServeMux, db *gorm.DB) {
	_, weaviateErr := vectorstore.GetWeaviateClient()
	h := &Handler{DB: db, VectorSyncEnabled: weaviateErr == nil}
	mux.HandleFunc("GET /components", h.GetComponents)
	mux.HandleFunc("POST /components", h.CreateComponent)
	mux.HandleFunc("GET /components/{id}", h.GetComponentByID)
	mux.HandleFunc("DELETE /components/{id}", h.DeleteComponent)
	mux.HandleFunc("PATCH /components/{id}", h.UpdateComponent)
}
