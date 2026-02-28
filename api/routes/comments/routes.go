package comments

import (
	"net/http"

	"gorm.io/gorm"
)

func RegisterCommentsRoutes(mux *http.ServeMux, db *gorm.DB) {
	// Comments routes are registered under /issues/{id}/comments in issue/routes.go
	// This function is kept for potential standalone comment routes in the future.
	_ = mux
	_ = db
}
