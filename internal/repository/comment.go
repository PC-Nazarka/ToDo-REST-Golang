package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"todo-list/internal/entity"
)

type CommentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (c *CommentRepository) Create(userId int, comment entity.CommentCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (text, user_id, post_id) VALUES ($1, $2, $3) RETURNING id;", commentTable)
	row := c.db.QueryRow(query, comment.Text, userId, comment.PostId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *CommentRepository) GetById(id int) (entity.Comment, error) {
	var comment entity.Comment
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", commentTable)
	if err := c.db.Get(&comment, query, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			return comment, errors.New("comment with this id not found")
		default:
			return comment, errors.New(fmt.Sprintf("error: %s", err.Error()))
		}
	}
	return comment, nil
}

func (c *CommentRepository) GetByPostId(id int) ([]entity.Comment, error) {
	var comments []entity.Comment
	query := fmt.Sprintf("SELECT * FROM %s WHERE post_id=$1 ORDER BY created_at DESC;", commentTable)
	if err := c.db.Select(&comments, query, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			return comments, errors.New("comments of post with this id not found")
		default:
			return comments, errors.New(fmt.Sprintf("error: %s", err.Error()))
		}
	}
	if comments == nil {
		comments = make([]entity.Comment, 0)
	}
	return comments, nil
}

func (c *CommentRepository) Update(id int, comment entity.CommentUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if comment.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, comment.Text)
		argId++
	}
	queryValues := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d;", commentTable, queryValues, argId)
	args = append(args, id)
	_, err := c.db.Exec(query, args...)
	return err
}

func (c *CommentRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", commentTable)
	_, err := c.db.Exec(query, id)
	return err
}
