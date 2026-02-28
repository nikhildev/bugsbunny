package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Change struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	OldValue  string    `json:"old_value" gorm:"not null"`
	NewValue  string    `json:"new_value" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" time_format:"2006-01-02 15:04:05" gorm:"autoCreateTime"`
}

func (c *Change) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		uuid, err := uuid.NewV7()
		if err != nil {
			return err
		}
		c.ID = uuid.String()
	}
	return nil
}
