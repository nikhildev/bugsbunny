package component

import "gorm.io/gorm"

// Handler holds dependencies for component HTTP handlers.
type Handler struct {
	DB *gorm.DB
}
