package repository

import "todo-app/internal/models"

type TaskRepository interface {
	Create(task models.Task) (*models.Task, error)
	GetByID(id string) (*models.Task, error)
	Update(task models.Task) (*models.Task, error)
	Delete(id string) error
	GetAll() ([]models.Task, error)
}
