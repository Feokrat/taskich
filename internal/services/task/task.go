package taskService

import (
	"github.com/Feokrat/taskich/internal/entities"
	"github.com/Feokrat/taskich/internal/models"
	taskRepository "github.com/Feokrat/taskich/internal/repositories/task"
)

type Task struct {
	repository *taskRepository.TaskRepository
}

type RepositoryInterface interface {
	Create(task models.Task) (int, error)
}

func New() *Task {
	return &Task{
		repository: taskRepository.New(),
	}
}

func (t *Task) Create(task entities.Task) (int, error) {
	model := models.Task{
		Name:        task.Name,
		Description: task.Description,
	}
	id, err := t.repository.Create(model)
	if err != nil {
		return 0, err
	}

	return id, nil
}
