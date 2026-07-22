package models

import "time"

type WeekParity int

const (
	WeekParityEvery WeekParity = 0 // каждая неделя
	WeekParityOdd   WeekParity = 1 // нечётная
	WeekParityEven  WeekParity = 2 // чётная
)

type ScheduleRule struct {
	ID          int        `json:"id"           db:"id"`
	UserID      int        `json:"user_id"      db:"user_id"`
	WorkTypeID  int        `json:"work_type_id" db:"work_type_id"`
	SubjectName string     `json:"subject_name" db:"subject_name"`
	DayOfWeek   int        `json:"day_of_week"  db:"day_of_week"`
	WeekParity  WeekParity `json:"week_parity"  db:"week_parity"`
	StartTime   TimeOfDay  `json:"start_time"   db:"start_time"`
	EndTime     TimeOfDay  `json:"end_time"     db:"end_time"`
	Room        string     `json:"room"         db:"room"`

	// Поля для безопасной замены расписания (без удаления истории)
	ValidFrom time.Time  `json:"valid_from" db:"valid_from"`
	ValidTo   *time.Time `json:"valid_to"   db:"valid_to"`
}

type ScheduleRuleGroups struct {
	ScheduleRuleID int `json:"schedule_rule_id" db:"schedule_rule_id"`
	GroupID        int `json:"group_id"         db:"group_id"`
}
