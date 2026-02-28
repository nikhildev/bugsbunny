package models

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
	ComponentID   string      `json:"component_id" gorm:"type:uuid; not null"`
	Component     Component   `json:"component" gorm:"foreignKey:ComponentID;references:ID"`
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
	BUG           IssueType = "bug"
	FEATURE       IssueType = "feature"
	SUPPORT       IssueType = "support"
	IMPROVEMENT   IssueType = "improvement"
	DOCUMENTATION IssueType = "documentation"
)

type IssueStatus string

const (
	NEW         IssueStatus = "new"
	RESOLVED    IssueStatus = "resolved"
	IN_PROGRESS IssueStatus = "in_progress"
	REOPENED    IssueStatus = "reopened"
	BLOCKED       IssueStatus = "blocked"
	ON_HOLD       IssueStatus = "on_hold"
	ISSUE_DELETED IssueStatus = "deleted"
)

type Severity string

const (
	LOW_SEVERITY      Severity = "low"
	MEDIUM_SEVERITY   Severity = "medium"
	HIGH_SEVERITY     Severity = "high"
	CRITICAL_SEVERITY Severity = "critical"
)

type Priority string

const (
	LOW_PRIORITY      Priority = "low"
	MEDIUM_PRIORITY   Priority = "medium"
	HIGH_PRIORITY     Priority = "high"
	CRITICAL_PRIORITY Priority = "critical"
)
