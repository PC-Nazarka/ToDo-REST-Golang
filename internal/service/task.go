package service

import (
	"errors"
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

func (t *TaskService) Update(userId, taskId int, task entity.TaskUpdate) error {
	taskExists, err := t.repo.GetById(taskId)
	if err != nil {
		return err
	}
	if userId != taskExists.UserId {
		return errors.New("you can't update task another user")
	}
	if err = task.Validate(); err != nil {
		return err
	}
	return t.repo.Update(taskId, task)
}

func (t *TaskService) Delete(userId, taskId int) error {
	task, err := t.repo.GetById(taskId)
	if err != nil {
		return err
	}
	if userId != task.UserId {
		return errors.New("you can't delete task another user")
	}
	return t.repo.Delete(taskId)
}

func (t *TaskService) GetByUserId(id int) ([]entity.Task, error) {
	return t.repo.GetByUserId(id)
}
