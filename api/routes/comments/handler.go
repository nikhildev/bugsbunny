package comments

import "gorm.io/gorm"

// Handler holds dependencies for comments HTTP handlers.
type Handler struct {
	DB *gorm.DB
}
