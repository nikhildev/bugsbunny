package repository

import (
	"github.com/nikhildev/bugsbunny/api/internal/model"
	"gorm.io/gorm"
)

type ProjectRepo interface {
	Create(project *model.Project) error
	GetByID(id string) (model.Project, error)
	GetAll() ([]model.Project, error)
	Update(id string, updates map[string]any) (int64, error)
	Delete(id string) error
}

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) Create(project *model.Project) error {
	return r.DB.Create(project).Error
}

func (r *ProjectRepository) GetByID(id string) (model.Project, error) {
	var project model.Project
	err := r.DB.First(&project, "id = ?", id).Error
	return project, err
}

func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	var projects []model.Project
	err := r.DB.Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) Update(id string, updates map[string]any) (int64, error) {
	result := r.DB.Model(&model.Project{}).Where("id = ?", id).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *ProjectRepository) Delete(id string) error {
	return r.DB.Model(&model.Project{}).Where("id = ?", id).Update("status", model.ProjectStatusDeleted).Error
}
