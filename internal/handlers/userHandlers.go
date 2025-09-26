package handlers

import (
	"PetProject/internal/userService"
	"PetProject/internal/web/users"
	"context"
)

type UserHandler struct {
	UserService userService.UserService
}

func NewUserHandler(userService userService.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) GetUsersUserIdTasks(_ context.Context, request users.GetUsersUserIdTasksRequestObject) (users.GetUsersUserIdTasksResponseObject, error) {
	tasks, err := h.UserService.GetTasksByUserID(request.UserId)
	if err != nil {
		return nil, err
	}
	response := users.GetUsersUserIdTasks200JSONResponse{}
	for _, tsk := range tasks {
		task := users.Task{
			Id:     &tsk.ID,
			Text:   &tsk.Text,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.UserService.CreateUser(userToCreate.Email, userToCreate.Password)

	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (h *UserHandler) PatchUsersUserId(_ context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {
	userRequest := request.Body
	userId := request.UserId
	updatedUser, err := h.UserService.UpdateUserByID(userId, userRequest.Email, userRequest.Password)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsersUserId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersUserId(_ context.Context, request users.DeleteUsersUserIdRequestObject) (users.DeleteUsersUserIdResponseObject, error) {
	userId := request.UserId
	err := h.UserService.DeleteUserByID(userId)
	if err != nil {
		return nil, err
	} else {
		response := users.DeleteUsersUserId204Response{}
		return response, nil
	}
}
