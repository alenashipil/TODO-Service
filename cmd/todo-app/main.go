package main

import (
	"todo-app/internal/handlers"
	"todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewMapTaskRepository()
	taskHandler := handlers.NewTaskHandler(repo)

	router := gin.Default()

	router.POST("/tasks", taskHandler.CreateTask)
	router.GET("/tasks", taskHandler.GetTasks)
	router.GET("/tasks/:id", taskHandler.GetTaskByID)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)

	router.Run(":8080")
}
