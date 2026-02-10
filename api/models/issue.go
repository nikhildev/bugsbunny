package models

type Issue struct {
	BaseModel
	Title         string      `json:"title" gorm:"size:255;not null"`
	Description   string      `json:"description" gorm:"type:text;not null"`
	Type          IssueType   `json:"type"`
	Status        IssueStatus `json:"status"`
	Assignee      *string     `json:"assignee" gorm:"size:64;not null"`
	Reporter      *string     `json:"reporter" gorm:"size:64;not null"`
	Components    []string    `json:"components" gorm:"not null"`
	Attachments   []string    `json:"attachments"`
	Tags          []string    `json:"tags"`
	Priority      *Priority   `json:"priority"`
	Severity      *Severity   `json:"severity"`
	Collaborators []string    `json:"collaborators"`
	CC            []string    `json:"cc"`
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
	BLOCKED     IssueStatus = "blocked"
	ON_HOLD     IssueStatus = "on_hold"
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
