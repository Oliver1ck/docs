package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	RoleID    int    `json:"role_id"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserLessons struct {
	UserID   int `json:"user_id"`
	LessonID int `json:"lesson_id"`
}
