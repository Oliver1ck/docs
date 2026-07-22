package models

import "time"

type User struct {
	ID        int       `json:"id"         db:"id"`
	Username  string    `json:"username"   db:"username"`
	Email     string    `json:"email"      db:"email"`
	Password  string    `json:"-"          db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	RoleID    int       `json:"role_id"    db:"role_id"`
}

type Role struct {
	ID   int    `json:"id"   db:"id"`
	Name string `json:"name" db:"name"`
}

type UserPermissions struct {
	UserID       int `json:"user_id"       db:"user_id"`
	PermissionID int `json:"permission_id" db:"permission_id"`
}
