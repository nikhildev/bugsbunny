package repository

import (
	"github.com/nikhildev/bugsbunny/api/internal/model"
	"gorm.io/gorm"
)

type CommentRepo interface {
	Create(comment *model.Comment) error
}

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) Create(comment *model.Comment) error {
	return r.DB.Create(comment).Error
}
