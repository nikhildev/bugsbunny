package project

import "gorm.io/gorm"

type Handler struct {
	DB                *gorm.DB
	VectorSyncEnabled bool
}
