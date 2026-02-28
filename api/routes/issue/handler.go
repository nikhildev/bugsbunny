package issue

import "gorm.io/gorm"

// Handler holds dependencies for issue HTTP handlers.
type Handler struct {
	DB *gorm.DB
}
