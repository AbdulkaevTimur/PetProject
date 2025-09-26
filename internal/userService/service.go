package userService

import (
	"PetProject/internal/taskService"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(email, password string) (*User, error)
	GetAllUsers() ([]*User, error)
	GetTasksByUserID(userId string) ([]*taskService.Task, error)
	GetUserByID(userId string) (*User, error)
	UpdateUserByID(userId string, email, password *string) (*User, error)
	DeleteUserByID(userId string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(email, password string) (*User, error) {
	user := &User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: password,
	}
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]*User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetTasksByUserID(userId string) ([]*taskService.Task, error) {
	if _, err := uuid.Parse(userId); err != nil {
		return nil, err
	}
	return s.repo.GetTasksByUserID(userId)
}

func (s *userService) GetUserByID(userId string) (*User, error) {
	if _, err := uuid.Parse(userId); err != nil {
		return nil, err
	}
	return s.repo.GetUserByID(userId)
}

func (s *userService) UpdateUserByID(userId string, email, password *string) (*User, error) {
	user, err := s.repo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}
	if email != nil {
		user.Email = *email
	}
	if password != nil {
		user.Password = *password
	}

	if err := s.repo.UpdateUserByID(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUserByID(userId string) error {
	if _, err := uuid.Parse(userId); err != nil {
		return err
	}
	return s.repo.DeleteUserByID(userId)
}
