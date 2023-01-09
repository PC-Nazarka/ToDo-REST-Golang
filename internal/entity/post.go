package entity

import (
	"errors"
	"time"
)

type Post struct {
	Id        int       `json:"id" db:"id"`
	Text      string    `json:"text" db:"text"`
	UserId    int       `json:"user_id" db:"user_id"`
	TaskId    int       `json:"task_id" db:"task_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type PostCreate struct {
	Text   string `json:"text" db:"text"`
	TaskId int    `json:"task_id" db:"task_id" binding:"required"`
}

type PostUpdate struct {
	Text *string `json:"text" db:"text"`
}

func (u *PostUpdate) Validate() error {
	if u.Text == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
