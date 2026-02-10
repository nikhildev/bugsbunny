package models

import "time"

type BaseModel struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at" time_format:"2006-01-02 15:04:05" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" time_format:"2006-01-02 15:04:05" gorm:"autoUpdateTime"`
}
