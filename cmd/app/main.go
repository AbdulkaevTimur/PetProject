package main

import (
	"PetProject/internal/db"
	"PetProject/internal/handlers"
	"PetProject/internal/taskService"
	"PetProject/internal/userService"
	"PetProject/internal/web/tasks"
	"PetProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	e := echo.New()

	tskRepo := taskService.NewTaskRepository(database)
	tskService := taskService.NewTaskService(tskRepo)
	tskHandlers := handlers.NewTaskHandler(tskService)

	usrRepo := userService.NewUserRepository(database)
	usrService := userService.NewUserService(usrRepo)
	usrHandlers := handlers.NewUserHandler(usrService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictTasksHandler := tasks.NewStrictHandler(tskHandlers, nil)
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(usrHandlers, nil)
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
