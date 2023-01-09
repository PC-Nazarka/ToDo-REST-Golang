package entity

import "errors"

type Task struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	IsDone      bool   `json:"is_done" db:"is_done"`
	UserId      int    `json:"user_id" db:"user_id"`
}

type TaskCreate struct {
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
	IsDone      bool   `json:"is_done" db:"is_done"`
}

type TaskUpdate struct {
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	IsDone      *bool   `json:"is_done" db:"is_done"`
}

func (t *TaskUpdate) Validate() error {
	if t.Description == nil && t.IsDone == nil && t.Name == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
