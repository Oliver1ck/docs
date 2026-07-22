package models

import "time"

// TrackStatus определяет статус записи о выполненной работе.
type TrackStatus string

const (
	TrackStatusPending   TrackStatus = "pending"
	TrackStatusConfirmed TrackStatus = "confirmed"
	TrackStatusRejected  TrackStatus = "rejected"
)

type Track struct {
	ID             int         `json:"id"               db:"id"`
	UserID         int         `json:"user_id"          db:"user_id"`
	ScheduleRuleID *int        `json:"schedule_rule_id" db:"schedule_rule_id"`
	WorkTypeID     int         `json:"work_type_id"     db:"work_type_id"`
	Date           time.Time   `json:"date"             db:"date"`
	AcademicHours  float64     `json:"academic_hours"   db:"academic_hours"`
	Status         TrackStatus `json:"status"           db:"status"`
	Comment        string      `json:"comment"          db:"comment"`
}

type TrackGroups struct {
	TrackID int `json:"track_id" db:"track_id"`
	GroupID int `json:"group_id" db:"group_id"`
}
