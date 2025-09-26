package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task *Task) error
	GetAllTasks() ([]*Task, error)
	GetTaskByID(taskId string) (*Task, error)
	UpdateTaskByID(task *Task) error
	DeleteTaskByID(taskId string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(text *Task) error {
	return r.db.Create(&text).Error
}

func (r *taskRepository) GetAllTasks() ([]*Task, error) {
	var tasks []*Task
	return tasks, r.db.Find(&tasks).Error
}

func (r *taskRepository) GetTaskByID(id string) (*Task, error) {
	var text *Task
	err := r.db.First(&text, "id = ?", id).Error
	return text, err
}

func (r *taskRepository) UpdateTaskByID(text *Task) error {
	return r.db.Save(&text).Error
}

func (r *taskRepository) DeleteTaskByID(taskId string) error {
	return r.db.Delete(&Task{}, "id = ?", taskId).Error
}
