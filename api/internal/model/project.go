package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	BaseModel
	Name            string        `json:"name" gorm:"not null"`
	ParentID        *string       `json:"parent_id"`
	Description     string        `json:"description" gorm:"not null"`
	Owner           string        `json:"owner" gorm:"index; not null"`
	Status          ProjectStatus `json:"status" gorm:"not null"`
	SlackChannelID  *string       `json:"slack_channel_id"`
	IsBotEnabled    bool          `json:"is_bot_enabled"`
	BotKnowledge    []string      `json:"bot_knowledge" gorm:"type:jsonb;serializer:json"`
	BotInstructions []string      `json:"bot_instructions" gorm:"type:jsonb;serializer:json"`
}

type ProjectStatus string

const (
	ProjectStatusActive   ProjectStatus = "active"
	ProjectStatusArchived ProjectStatus = "archived"
	ProjectStatusDeleted  ProjectStatus = "deleted"
)

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}
		p.ID = uuid.String()
	}
	return nil
}
