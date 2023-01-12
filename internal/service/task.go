package service

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
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

func (t *TaskService) CreateBulk(userId int, tasks []entity.TaskCreate) ([]int, error) {
	createdTasksIds := make([]int, 0)
	for _, task := range tasks {
		id, err := t.repo.Create(userId, task)
		if err != nil {
			return make([]int, 0), err
		}
		createdTasksIds = append(createdTasksIds, id)
	}
	return createdTasksIds, nil
}

func (t *TaskService) GetById(id int) (entity.Task, error) {
	return t.repo.GetById(id)
}

func (t *TaskService) GetByIds(ids []int) ([]entity.Task, error) {
	createdTasks := make([]entity.Task, 0)
	for _, id := range ids {
		task, err := t.repo.GetById(id)
		if err != nil {
			return make([]entity.Task, 0), err
		}
		createdTasks = append(createdTasks, task)
	}
	return createdTasks, nil
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

func (t *TaskService) ParseFile(path string) ([]entity.TaskCreate, error) {
	var tasks = make([]entity.TaskCreate, 0)
	file, err := os.Open(path)
	if err != nil {
		return tasks, nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	if _, err := reader.Read(); err != nil {
		return tasks, nil
	}
	records, err := reader.ReadAll()
	if err != nil {
		return tasks, nil
	}
	for _, task := range records {
		isDone, err := strconv.ParseBool(task[2])
		if err != nil {
			return tasks, err
		}
		tempTask := entity.TaskCreate{
			Name:        task[0],
			Description: task[1],
			IsDone:      isDone,
		}
		tasks = append(tasks, tempTask)
	}
	if tasks == nil {
		tasks = make([]entity.TaskCreate, 0)
	}
	return tasks, nil
}

func (t *TaskService) WriteFile(tasks []entity.Task, path string) error {
	headers := []string{
		"Name",
		"Description",
		"IsDone",
	}
	records := make([][]string, 0)
	for _, task := range tasks {
		temp := make([]string, 0)
		temp = append(temp, task.Name)
		temp = append(temp, task.Description)
		temp = append(temp, strconv.FormatBool(task.IsDone))
		records = append(records, temp)
	}
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Comma = ';'
	if err := writer.Write(headers); err != nil {
		return err
	}
	err = writer.WriteAll(records)
	return err
}
