package taskService

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	uuid1 = "550e8400-e29b-41d4-a716-446655440000"
	uuid2 = "550e8400-e29b-41d4-a716-446655440001"
	uuid3 = "550e8400-e29b-41d4-a716-446655440002"
)

func TestTaskService_CreateTask(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name: "успешное создание задачи",
			mockSetup: func(m *MockTaskRepository) {
				m.On("CreateTask", mock.AnythingOfType("*taskService.Task")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при создании",
			mockSetup: func(m *MockTaskRepository) {
				m.On("CreateTask", mock.AnythingOfType("*taskService.Task")).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)
			task, err := service.CreateTask(uuid1, "Test task")

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, task)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, task)
				assert.Equal(t, "Test task", task.Text)
				assert.Equal(t, uuid1, task.UserID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestTaskService_GetAllTasks(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name: "успешное получение всех задач",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTasks").Return([]*Task{
					{ID: uuid1, Text: "Task 1"},
					{ID: uuid2, Text: "Task 2"},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при получении задач",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTasks").Return([]*Task(nil), errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)
			tasks, err := service.GetAllTasks()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, tasks)
			} else {
				assert.NoError(t, err)
				assert.Len(t, tasks, 2)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestTaskService_GetTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(m *MockTaskRepository, id string)
		wantErr   bool
	}{
		{
			name: "задача найдена",
			id:   uuid1,
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("GetTaskByID", id).Return(&Task{ID: id, Text: "Task 1"}, nil)
			},
			wantErr: false,
		},
		{
			name: "задача не найдена",
			id:   uuid3,
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("GetTaskByID", id).Return((*Task)(nil), errors.New("not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id)

			service := NewTaskService(mockRepo)
			task, err := service.GetTaskByID(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, task)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.id, task.ID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestTaskService_UpdateTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		task      *Task
		mockSetup func(m *MockTaskRepository, task *Task)
		wantErr   bool
	}{
		{
			name: "успешное обновление",
			task: &Task{ID: uuid1, Text: "Updated Task"},
			mockSetup: func(m *MockTaskRepository, task *Task) {
				m.On("GetTaskByID", task.ID).Return(&Task{ID: task.ID, Text: "Old Task"}, nil)
				m.On("UpdateTaskByID", mock.AnythingOfType("*taskService.Task")).Return(task, nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при обновлении",
			task: &Task{ID: uuid2, Text: "Fail Task"},
			mockSetup: func(m *MockTaskRepository, task *Task) {
				m.On("GetTaskByID", task.ID).Return(&Task{ID: task.ID, Text: "Old Task"}, nil)
				m.On("UpdateTaskByID", mock.AnythingOfType("*taskService.Task")).Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.task)

			service := NewTaskService(mockRepo)
			updated, err := service.UpdateTaskByID(tt.task.ID, &tt.task.Text, &tt.task.IsDone)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, updated)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.task.Text, updated.Text)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestTaskService_DeleteTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(m *MockTaskRepository, id string)
		wantErr   bool
	}{
		{
			name: "успешное удаление",
			id:   uuid1,
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("DeleteTaskByID", id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при удалении",
			id:   uuid2,
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("DeleteTaskByID", id).Return(errors.New("delete error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id)

			service := NewTaskService(mockRepo)
			err := service.DeleteTaskByID(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
