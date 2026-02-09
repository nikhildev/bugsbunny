package models

import "time"

type Component struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	ParentID        *int            `json:"parent_id"`
	CreatedAt       time.Time       `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt       time.Time       `json:"updated_at" time_format:"2006-01-02 15:04:05"`
	Description     string          `json:"description"`
	Owner           string          `json:"owner"`
	Status          ComponentStatus `json:"status"`
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
