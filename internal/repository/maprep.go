package repository

import (
	"errors"
	"sync"
	"todo-app/internal/models"

	"github.com/google/uuid"
)

type MapTaskRepository struct {
	sync.Mutex
	tasks map[string]models.Task
}

func NewMapTaskRepository() *MapTaskRepository {
	return &MapTaskRepository{
		tasks: make(map[string]models.Task),
	}
}
func (r *MapTaskRepository) Create(task models.Task) (*models.Task, error) {

	r.Lock()
	defer r.Unlock()
	task.ID = uuid.New().String()
	r.tasks[task.ID] = task
	return &task, nil
}

func (r *MapTaskRepository) GetByID(id string) (*models.Task, error) {

	r.Lock()
	defer r.Unlock()
	task, exists := r.tasks[id]

	if !exists {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

func (r *MapTaskRepository) Delete(id string) error {
	r.Lock()
	defer r.Unlock()
	if _, exists := r.tasks[id]; !exists {
		return errors.New("task is not found")
	}
	delete(r.tasks, id)
	return nil
}

func (r *MapTaskRepository) Update(task models.Task) (*models.Task, error) {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return nil, errors.New("task not found")
	}
	r.tasks[task.ID] = task
	return &task, nil
}

func (r *MapTaskRepository) GetAll() ([]models.Task, error) {
	r.Lock()
	defer r.Unlock()

	tasks := make([]models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
