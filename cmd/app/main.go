package main

import (
	"log"

	"github.com/Feokrat/taskich/internal/entities"
	taskService "github.com/Feokrat/taskich/internal/services/task"
)

func main() {
	service := taskService.New()

	tmp := entities.Task{
		Name:        "New",
		Description: "Test",
	}
	id, err := service.Create(tmp)
	if err != nil {
		log.Fatalf("coul not create task %v", tmp)
	}

	log.Printf("created task with id %d", id)
}
