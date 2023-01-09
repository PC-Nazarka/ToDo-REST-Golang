package service

import (
	"todo-list/internal/entity"
	"todo-list/internal/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) Create(userId int, task entity.TaskCreate) (int, error) {
	return t.repo.Create(userId, task)
}

func (t *TaskService) GetById(id int) (entity.Task, error) {
	return t.repo.GetById(id)
}

func (t *TaskService) Update(id int, task entity.TaskUpdate) error {
	if err := task.Validate(); err != nil {
		return err
	}
	return t.repo.Update(id, task)
}

func (t *TaskService) Delete(id int) error {
	return t.repo.Delete(id)
}

func (t *TaskService) GetByUserId(id int) ([]entity.Task, error) {
	return t.repo.GetByUserId(id)
}
