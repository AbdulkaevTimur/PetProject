package userService

import (
	"PetProject/internal/taskService"
	"time"
)

type User struct {
	ID        string             `json:"id"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Tasks     []taskService.Task `json:"tasks"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt *time.Time         `sql:"index" json:"deleted_at"`
}
