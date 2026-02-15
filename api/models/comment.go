package models

type Comment struct {
	BaseModel
	IssueID     string   `json:"issue_id" gorm:"type:uuid;not null"`
	Content     string   `json:"content" gorm:"type:text;not null"`
	Author      string   `json:"author" gorm:"not null"`
	Attachments []string `json:"attachments" gorm:"type:jsonb;serializer:json"`
}
