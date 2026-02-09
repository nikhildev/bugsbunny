package models

import "time"

type Issue struct {
	ID            int         `json:"id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	Type          IssueType   `json:"type"`
	Status        IssueStatus `json:"status"`
	CreatedAt     time.Time   `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt     time.Time   `json:"updated_at" time_format:"2006-01-02 15:04:05"`
	Assignee      *string     `json:"assignee"`
	Reporter      *string     `json:"reporter"`
	Components    []string    `json:"components"`
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
