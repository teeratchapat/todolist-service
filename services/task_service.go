package services

import (
	"errors"
	"todolist-service/models"
	"todolist-service/repositories"
)

type TaskService struct {
	Repo *repositories.TaskRepository
}

func (s *TaskService) CreateTask(task *models.Task) error {
	if task.Detail == "" || len(task.Detail) > 50 {
		return errors.New("Detail must not be empty and have 50 characters")
	}
	return s.Repo.CreateTask(task)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.Repo.GetTasks()
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	return s.Repo.GetTaskByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.Repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id int) error {
	task, err := s.Repo.GetTaskByID(id)
	if err != nil || task == nil {
		return errors.New("task not found")
	}

	return s.Repo.DeleteTask(id)
}
