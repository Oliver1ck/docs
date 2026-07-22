package models

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// TimeOfDay представляет время суток в формате "HH:MM".
// Сканируется из PostgreSQL TIME, сериализуется в JSON как строка "08:30".
type TimeOfDay string

// Scan реализует интерфейс sql.Scanner — позволяет pgx сканировать TIME в TimeOfDay.
// pgx возвращает TIME в виде строки "08:30:00", мы обрезаем секунды.
func (t *TimeOfDay) Scan(src any) error {
	if src == nil {
		*t = ""
		return nil
	}
	var s string
	switch v := src.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return fmt.Errorf("TimeOfDay: cannot scan type %T", src)
	}
	// PostgreSQL TIME может вернуть "08:30:00" — обрезаем секунды
	parts := strings.SplitN(s, ":", 3)
	if len(parts) < 2 {
		return fmt.Errorf("TimeOfDay: invalid format %q", s)
	}
	cleaned := parts[0] + ":" + parts[1]
	if err := validateTimeOfDay(cleaned); err != nil {
		return err
	}
	*t = TimeOfDay(cleaned)
	return nil
}

// Value реализует интерфейс driver.Valuer — позволяет pgx записывать TimeOfDay в TIME.
func (t TimeOfDay) Value() (driver.Value, error) {
	if t == "" {
		return nil, nil
	}
	if err := validateTimeOfDay(string(t)); err != nil {
		return nil, err
	}
	return string(t), nil
}

func validateTimeOfDay(s string) error {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("TimeOfDay: expected HH:MM, got %q", s)
	}
	h, err := strconv.Atoi(parts[0])
	if err != nil || h < 0 || h > 23 {
		return fmt.Errorf("TimeOfDay: invalid hours in %q", s)
	}
	m, err := strconv.Atoi(parts[1])
	if err != nil || m < 0 || m > 59 {
		return fmt.Errorf("TimeOfDay: invalid minutes in %q", s)
	}
	return nil
}
