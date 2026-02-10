package models

type Comment struct {
	BaseModel
	IssueID     string   `json:"issue_id" gorm:"not null"`
	Content     string   `json:"content" gorm:"not null"`
	Author      string   `json:"author" gorm:"not null"`
	Attachments []string `json:"attachments" gorm:"not null"`
}
