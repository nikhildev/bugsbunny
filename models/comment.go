package models

import "time"

type Comment struct {
	ID          int       `json:"id"`
	IssueID     string    `json:"issue_id"`
	CreatedAt   time.Time `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `json:"updated_at" time_format:"2006-01-02 15:04:05"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	Attachments []string  `json:"attachments"`
}
