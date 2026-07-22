package models

type WorkCategory struct {
	ID        int    `json:"id"         db:"id"`
	Name      string `json:"name"       db:"name"`
	SortOrder int    `json:"sort_order" db:"sort_order"`
}

type WorkType struct {
	ID         int    `json:"id"          db:"id"`
	Name       string `json:"name"        db:"name"`
	ShortName  string `json:"short_name"  db:"short_name"`
	CategoryID int    `json:"category_id" db:"category_id"`
	IsActive   bool   `json:"is_active"   db:"is_active"`
}
