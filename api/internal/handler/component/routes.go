package component

import (
	"net/http"

	"gorm.io/gorm"
)

func RegisterComponentRoutes(mux *http.ServeMux, db *gorm.DB) {
	h := &Handler{DB: db}
	mux.HandleFunc("GET /components", h.GetComponents)
	mux.HandleFunc("POST /components", h.CreateComponent)
	mux.HandleFunc("GET /components/{id}", h.GetComponentByID)
	mux.HandleFunc("DELETE /components/{id}", h.DeleteComponent)
	mux.HandleFunc("PATCH /components/{id}", h.UpdateComponent)
}
