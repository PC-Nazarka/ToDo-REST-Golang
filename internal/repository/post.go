package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"todo-list/internal/entity"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (p *PostRepository) Create(userId int, post entity.PostCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (text, user_id, task_id) VALUES ($1, $2, $3) RETURNING id;", postTable)
	row := p.db.QueryRow(query, post.Text, userId, post.TaskId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostRepository) GetById(id int) (entity.Post, error) {
	var post entity.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", postTable)
	if err := p.db.Get(&post, query, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			return post, errors.New("post with this id not found")
		default:
			return post, errors.New(fmt.Sprintf("error: %s", err.Error()))
		}
	}
	return post, nil
}

func (p *PostRepository) GetAll() ([]entity.Post, error) {
	var posts []entity.Post
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY created_at DESC;", postTable)
	if err := p.db.Select(&posts, query); err != nil {
		switch err {
		case sql.ErrNoRows:
			return posts, errors.New("posts not found")
		default:
			return posts, errors.New(fmt.Sprintf("error: %s", err.Error()))
		}
	}
	return posts, nil
}

func (p *PostRepository) Update(id int, post entity.PostUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if post.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, post.Text)
		argId++
	}
	queryValues := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d;", postTable, queryValues, argId)
	args = append(args, id)
	_, err := p.db.Exec(query, args...)
	return err
}

func (p *PostRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", postTable)
	_, err := p.db.Exec(query, id)
	return err
}

func (p *PostRepository) GetByUserId(id int) ([]entity.Post, error) {
	var posts []entity.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 ORDER BY created_at DESC;", postTable)
	if err := p.db.Select(&posts, query, id); err != nil {
		return posts, err
	}
	return posts, nil
}
