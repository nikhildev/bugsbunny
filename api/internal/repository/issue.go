package repository

import (
	"github.com/nikhildev/bugsbunny/api/internal/model"
	"gorm.io/gorm"
)

type IssueRepo interface {
	Create(issue *model.Issue) error
	GetByID(id string) (model.Issue, error)
	GetAll() ([]model.Issue, error)
	Update(id string, updates map[string]any) (int64, error)
	Delete(id string) error
}

type IssueRepository struct {
	DB *gorm.DB
}

func NewIssueRepository(db *gorm.DB) *IssueRepository {
	return &IssueRepository{DB: db}
}

func (r *IssueRepository) Create(issue *model.Issue) error {
	return r.DB.Create(issue).Error
}

func (r *IssueRepository) GetByID(id string) (model.Issue, error) {
	var issue model.Issue
	err := r.DB.First(&issue, "id = ?", id).Error
	return issue, err
}

func (r *IssueRepository) GetAll() ([]model.Issue, error) {
	var issues []model.Issue
	err := r.DB.Preload("Reporter").Preload("Assignee").Preload("Project").Preload("Collaborators").Preload("CC").Find(&issues).Error
	return issues, err
}

func (r *IssueRepository) Update(id string, updates map[string]any) (int64, error) {
	result := r.DB.Model(&model.Issue{}).Where("id = ?", id).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *IssueRepository) Delete(id string) error {
	return r.DB.Model(&model.Issue{}).Where("id = ?", id).Update("status", model.IssueStatusDeleted).Error
}
