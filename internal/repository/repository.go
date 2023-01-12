package repository

import (
	"todo-list/internal/entity"

	"github.com/jmoiron/sqlx"
)

const (
	userTable    = "users"
	taskTable    = "tasks"
	postTable    = "posts"
	commentTable = "comments"
	roomTable    = "rooms"
)

type User interface {
	GetByUsernameAndPassword(username, password string) (int, error)
	GetById(id int) (entity.User, error)
	GetAll() ([]entity.User, error)
	Create(user entity.UserCreate) (int, error)
	Update(id int, user entity.UserUpdate) error
	Delete(id int) error
}

type Task interface {
	Create(userId int, task entity.TaskCreate) (int, error)
	GetById(id int) (entity.Task, error)
	Update(id int, task entity.TaskUpdate) error
	Delete(id int) error
	GetByUserId(id int) ([]entity.Task, error)
}

type Post interface {
	Create(userId int, post entity.PostCreate) (int, error)
	GetById(id int) (entity.Post, error)
	GetAll() ([]entity.Post, error)
	Update(id int, post entity.PostUpdate) error
	Delete(id int) error
	GetByUserId(id int) ([]entity.Post, error)
}

type Comment interface {
	Create(userId int, comment entity.CommentCreate) (int, error)
	GetById(id int) (entity.Comment, error)
	GetByPostId(id int) ([]entity.Comment, error)
	Update(id int, comment entity.CommentUpdate) error
	Delete(id int) error
}

type Repository struct {
	User
	Task
	Post
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Task:    NewTaskRepository(db),
		Post:    NewPostRepository(db),
		Comment: NewCommentRepository(db),
	}
}
