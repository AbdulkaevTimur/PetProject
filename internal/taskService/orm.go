package taskService

import (
	"time"
)

type Task struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	Text      string     `json:"text"`
	IsDone    bool       `json:"is_done"`
	UserID    string     `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
