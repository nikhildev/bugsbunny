package project

import (
	"github.com/nikhildev/bugsbunny/api/internal/vectorstore"
	"gorm.io/gorm"
)

type Handler struct {
	DB          *gorm.DB
	VectorStore *vectorstore.VectorStore
}
