package entity

import (
	"errors"
	"time"
)

type Comment struct {
	Id        int       `json:"id" db:"id"`
	Text      string    `json:"text" db:"text"`
	UserId    int       `json:"user_id" db:"user_id"`
	PostId    int       `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CommentCreate struct {
	Text   string `json:"text" db:"text" binding:"required"`
	PostId int    `json:"post_id" db:"post_id" binding:"required"`
}

type CommentUpdate struct {
	Text *string `json:"text" db:"text"`
}

func (u *CommentUpdate) Validate() error {
	if u.Text == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
