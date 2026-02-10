package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at" time_format:"2006-01-02 15:04:05"`
	Role      UserRole  `json:"role"`
}

type UserRole string

const (
	Admin  UserRole = "admin"
	Viewer UserRole = "viewer"
	Editor UserRole = "editor"
)
