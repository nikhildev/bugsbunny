package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Component struct {
	BaseModel
	Name            string          `json:"name" gorm:"not null"`
	ParentID        *string         `json:"parent_id"`
	Description     string          `json:"description" gorm:"not null"`
	Owner           string          `json:"owner" gorm:"index; not null"`
	Status          ComponentStatus `json:"status" gorm:"not null"`
	SlackChannelID  *string         `json:"slack_channel_id"`
	IsBotEnabled    bool            `json:"is_bot_enabled"`
	BotKnowledge    []string        `json:"bot_knowledge" gorm:"type:jsonb;serializer:json"`
	BotInstructions []string        `json:"bot_instructions" gorm:"type:jsonb;serializer:json"`
}

type ComponentStatus string

const (
	ACTIVE   ComponentStatus = "active"
	ARCHIVED ComponentStatus = "archived"
)

func (c *Component) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewV7()

	if err != nil {
		return err
	}

	c.ID = uuid.String()
	return nil
}
