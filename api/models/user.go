package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username  string     `json:"username" gorm:"size:64;not null;unique"`
	Email     string     `json:"email" gorm:"not null;unique"`
	Password  string     `json:"password" gorm:"size:64;not null"`
	Role      UserRole   `json:"role" gorm:"not null"`
	IsActive  bool       `json:"is_active" gorm:"not null"`
	IsDeleted bool       `json:"is_deleted" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted_at" time_format:"2006-01-02 15:04:05" gorm:"index"`
}

type UserRole string

const (
	Admin  UserRole = "admin"
	Viewer UserRole = "viewer"
	Editor UserRole = "editor"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.ID = uuid.String()
	return nil
}
