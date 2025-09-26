package userService

import (
	"PetProject/internal/taskService"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetAllUsers() ([]*User, error) {
	args := m.Called()
	return args.Get(0).([]*User), args.Error(1)
}

func (m *MockUserRepository) GetTasksByUserID(userId string) ([]*taskService.Task, error) {
	args := m.Called(userId)
	return args.Get(0).([]*taskService.Task), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(userId string) (*User, error) {
	args := m.Called(userId)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserByID(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUserByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
