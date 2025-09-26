package userService

import (
	"PetProject/internal/taskService"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *User) error
	GetAllUsers() ([]*User, error)
	GetTasksByUserID(userId string) ([]*taskService.Task, error)
	GetUserByID(id string) (*User, error)
	UpdateUserByID(user *User) error
	DeleteUserByID(userId string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetAllUsers() ([]*User, error) {
	var users []*User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetTasksByUserID(userId string) ([]*taskService.Task, error) {
	var tasks []*taskService.Task
	err := r.db.Where("user_id = ?", userId).Find(&tasks).Error
	return tasks, err
}

func (r *userRepository) GetUserByID(id string) (*User, error) {
	var user *User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateUserByID(user *User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUserByID(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
