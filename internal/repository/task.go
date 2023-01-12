package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"todo-list/internal/entity"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Create(userId int, task entity.TaskCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, is_done, user_id) VALUES ($1, $2, $3, $4) RETURNING id;", taskTable)
	row := t.db.QueryRow(query, task.Name, task.Description, task.IsDone, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TaskRepository) GetById(id int) (entity.Task, error) {
	var task entity.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", taskTable)
	if err := t.db.Get(&task, query, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			return task, errors.New("task with this id not found")
		default:
			return task, errors.New(fmt.Sprintf("error: %s", err.Error()))
		}
	}
	return task, nil
}

func (t *TaskRepository) Update(id int, task entity.TaskUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if task.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, task.Name)
		argId++
	}
	if task.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, task.Description)
		argId++
	}
	if task.IsDone != nil {
		setValues = append(setValues, fmt.Sprintf("is_done=$%d", argId))
		args = append(args, task.IsDone)
		argId++
	}
	queryValues := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d;", taskTable, queryValues, argId)
	args = append(args, id)
	_, err := t.db.Exec(query, args...)
	return err
}

func (t *TaskRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", taskTable)
	_, err := t.db.Exec(query, id)
	return err
}

func (t *TaskRepository) GetByUserId(id int) ([]entity.Task, error) {
	var tasks []entity.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 ORDER BY id;", taskTable)
	if err := t.db.Select(&tasks, query, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			return tasks, errors.New("tasks of user with id not found")
		default:
			return tasks, errors.New(fmt.Sprintf("error: %s", err.Error()))
		}
	}
	if tasks == nil {
		tasks = make([]entity.Task, 0)
	}
	return tasks, nil
}
