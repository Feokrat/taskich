package taskRepository

import (
	"log"
	"math/rand"
	"time"

	"github.com/Feokrat/taskich/internal/models"
)

type TaskRepository struct {
}

func New() *TaskRepository {
	return &TaskRepository{}
}

func (*TaskRepository) Create(task models.Task) (int, error) {
	rand.Seed(time.Now().Unix())
	id := rand.Intn(10)
	log.Printf("Created task with name: %v, description: %v and id: %d", task.Name, task.Description, id)
	return id, nil
}
