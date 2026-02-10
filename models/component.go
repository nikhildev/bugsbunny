package models

type Component struct {
	BaseModel
	Name            string          `json:"name" gorm:"not null"`
	ParentID        *string         `json:"parent_id"`
	Description     string          `json:"description" gorm:"not null"`
	Owner           string          `json:"owner" gorm:"not null"`
	Status          ComponentStatus `json:"status" gorm:"not null"`
	SlackChannelID  *string         `json:"slack_channel_id"`
	IsBotEnabled    bool            `json:"is_bot_enabled"`
	BotKnowledge    []string        `json:"bot_knowledge"`
	BotInstructions []string        `json:"bot_instructions"`
}

type ComponentStatus string

const (
	ACTIVE   ComponentStatus = "active"
	ARCHIVED ComponentStatus = "archived"
)
