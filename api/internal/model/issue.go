package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Issue struct {
	BaseModel
	Title         string      `json:"title" gorm:"size:255;not null"`
	Description   string      `json:"description" gorm:"type:text;not null"`
	Type          IssueType   `json:"type"`
	Status        IssueStatus `json:"status" gorm:"not null;index;default:new"`
	AssigneeId    *string     `json:"assignee_id" gorm:"type:uuid;index"`
	Assignee      User        `json:"assignee" gorm:"foreignKey:AssigneeId;references:ID"`
	ReporterId    string      `json:"reporter_id" gorm:"type:uuid; not null"`
	Reporter      User        `json:"reporter" gorm:"foreignKey:ReporterId;references:ID"`
	ProjectID     string      `json:"project_id" gorm:"type:uuid; not null"`
	Project       Project     `json:"project" gorm:"foreignKey:ProjectID;references:ID"`
	Attachments   []string    `json:"attachments" gorm:"type:jsonb;serializer:json"`
	Priority      Priority    `json:"priority" gorm:"not null;index;default:low"`
	Severity      Severity    `json:"severity" gorm:"not null;index;default:low"`
	Collaborators []User      `json:"collaborators" gorm:"many2many:issue_collaborators;foreignKey:ID;references:ID"`
	CC            []User      `json:"cc" gorm:"many2many:issue_cc;foreignKey:ID;references:ID"`
	Comments      []Comment   `json:"comments" gorm:"foreignKey:IssueID;references:ID"`
}

func (i *Issue) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}
		i.ID = uuid.String()
	}
	return nil
}

type IssueType string

const (
	IssueTypeBug           IssueType = "bug"
	IssueTypeFeature       IssueType = "feature"
	IssueTypeSupport       IssueType = "support"
	IssueTypeImprovement   IssueType = "improvement"
	IssueTypeDocumentation IssueType = "documentation"
)

type IssueStatus string

const (
	IssueStatusNew       IssueStatus = "new"
	IssueStatusResolved  IssueStatus = "resolved"
	IssueStatusProgress  IssueStatus = "in_progress"
	IssueStatusReopened  IssueStatus = "reopened"
	IssueStatusBlocked   IssueStatus = "blocked"
	IssueStatusOnHold    IssueStatus = "on_hold"
	IssueStatusDeleted   IssueStatus = "deleted"
)

type Severity string

const (
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

type Priority string

const (
	PriorityLow      Priority = "low"
	PriorityMedium   Priority = "medium"
	PriorityHigh     Priority = "high"
	PriorityCritical Priority = "critical"
)
