package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	BaseModel
	IssueID     string    `json:"issue_id" gorm:"type:uuid;not null;index"`
	Content     string    `json:"content" gorm:"type:text;not null"`
	Author      uuid.UUID `json:"author" gorm:"type:uuid;not null;index"`
	Attachments []string  `json:"attachments" gorm:"type:jsonb;serializer:json"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}
		c.ID = id.String()
	}
	return nil
}
