package taskService

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(userId, text string) (*Task, error)
	GetAllTasks() ([]*Task, error)
	GetTaskByID(taskId string) (*Task, error)
	UpdateTaskByID(taskId string, text *string, isDone *bool) (*Task, error)
	DeleteTaskByID(taskId string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(userId, text string) (*Task, error) {
	task := &Task{
		ID:     uuid.NewString(),
		Text:   text,
		UserID: userId,
	}
	if err := s.repo.CreateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) GetAllTasks() ([]*Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(taskId string) (*Task, error) {
	if _, err := uuid.Parse(taskId); err != nil {
		return nil, err
	}
	return s.repo.GetTaskByID(taskId)
}

func (s *taskService) UpdateTaskByID(taskId string, text *string, isDone *bool) (*Task, error) {
	task, err := s.repo.GetTaskByID(taskId)
	if err != nil {
		return nil, err
	}
	if text != nil {
		task.Text = *text
	}
	if isDone != nil {
		task.IsDone = *isDone
	}
	if err := s.repo.UpdateTaskByID(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) DeleteTaskByID(taskId string) error {
	if _, err := uuid.Parse(taskId); err != nil {
		return err
	}
	return s.repo.DeleteTaskByID(taskId)
}
