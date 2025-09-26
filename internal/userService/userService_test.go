package userService

import (
	"PetProject/internal/taskService"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	mockRepo.On("CreateUser", mock.AnythingOfType("*userService.User")).Return(nil)

	email := "test@example.com"
	pass := "pass123"

	createdUser, err := svc.CreateUser(email, pass)
	assert.NoError(t, err)
	assert.Equal(t, email, createdUser.Email)
	assert.Equal(t, pass, createdUser.Password)
	mockRepo.AssertCalled(t, "CreateUser", mock.AnythingOfType("*userService.User"))
}

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	users := []*User{
		{ID: uuid.NewString(), Email: "a@test.com", Password: "123"},
		{ID: uuid.NewString(), Email: "b@test.com", Password: "456"},
	}

	mockRepo.On("GetAllUsers").Return(users, nil)

	result, err := svc.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, users, result)
}

func TestUserService_GetTasksByUserId(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	id := uuid.NewString()

	tasks := []*taskService.Task{
		{ID: "task1", Text: "Task 1"},
		{ID: "task2", Text: "Task 2"},
	}

	mockRepo.On("GetTasksByUserID", mock.Anything).Return(tasks, nil)
	result, err := svc.GetTasksByUserID(id)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, tasks, result)
}

func TestUserService_GetUserById(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	id := uuid.NewString()
	user := &User{ID: id, Email: "a@test.com", Password: "123"}

	mockRepo.On("GetUserByID", id).Return(user, nil)

	result, err := svc.GetUserByID(id)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestUserService_UpdateUserById(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	id := uuid.NewString()
	user := &User{ID: id, Email: "old@test.com", Password: "oldpass"}

	mockRepo.On("GetUserByID", id).Return(user, nil)
	mockRepo.On("UpdateUserByID", mock.AnythingOfType("*userService.User")).Return(nil)

	newEmail := "new@test.com"
	newPass := "newpass"

	updatedUser, err := svc.UpdateUserByID(id, &newEmail, &newPass)
	assert.NoError(t, err)
	assert.Equal(t, newEmail, updatedUser.Email)
	assert.Equal(t, newPass, updatedUser.Password)
}

func TestUserService_DeleteUserById(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	id := uuid.NewString()

	mockRepo.On("DeleteUserByID", id).Return(nil)

	err := svc.DeleteUserByID(id)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteUserByID", id)
}

func TestUserService_UpdateUserById_PartialNil(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo)

	id := uuid.NewString()
	user := &User{ID: id, Email: "old@test.com", Password: "oldpass"}

	mockRepo.On("GetUserByID", id).Return(user, nil)
	mockRepo.On("UpdateUserByID", mock.AnythingOfType("*userService.User")).Return(nil)

	// Обновляем только Email
	newEmail := "new@test.com"

	updatedUser, err := svc.UpdateUserByID(id, &newEmail, nil)
	assert.NoError(t, err)
	assert.Equal(t, newEmail, updatedUser.Email)
	assert.Equal(t, "oldpass", updatedUser.Password)
}
