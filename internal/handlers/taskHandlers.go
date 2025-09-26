package handlers

import (
	"PetProject/internal/taskService"
	"PetProject/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	TaskService taskService.TaskService
}

func NewTaskHandler(taskService taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.TaskService.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Text:   &tsk.Text,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasksUserId(_ context.Context, request tasks.PostTasksUserIdRequestObject) (tasks.PostTasksUserIdResponseObject, error) {
	taskRequest := request.Body
	userId := request.UserId
	taskToCreate := taskService.Task{
		Text: *taskRequest.Text,
	}
	createdTask, err := h.TaskService.CreateTask(userId, taskToCreate.Text)

	if err != nil {
		return nil, err
	}
	response := tasks.PostTasksUserId201JSONResponse{
		Id:     &createdTask.ID,
		Text:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil
}
func (h *TaskHandler) PatchTasksTaskId(_ context.Context, request tasks.PatchTasksTaskIdRequestObject) (tasks.PatchTasksTaskIdResponseObject, error) {
	taskRequest := request.Body
	taskId := request.TaskId
	updatedTask, err := h.TaskService.UpdateTaskByID(taskId, taskRequest.Text, taskRequest.IsDone)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasksTaskId200JSONResponse{
		Id:     &updatedTask.ID,
		Text:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
		UserId: &updatedTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksTaskId(_ context.Context, request tasks.DeleteTasksTaskIdRequestObject) (tasks.DeleteTasksTaskIdResponseObject, error) {
	taskId := request.TaskId
	err := h.TaskService.DeleteTaskByID(taskId)
	if err != nil {
		return nil, err
	} else {
		response := tasks.DeleteTasksTaskId204Response{}
		return response, nil
	}
}
