package service

import (
	"todo-list/internal/entity"
	"todo-list/internal/repository"
)

type JWTAuthorization interface {
	GenerateToken(id int) (string, error)
	ParseToken(accessToken string) (int, error)
}

type User interface {
	GetByUsernameAndPassword(username, password string) (int, error)
	GetById(id int) (entity.User, error)
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
	Create(userId int, task entity.PostCreate) (int, error)
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

type Service struct {
	JWTAuthorization
	User
	Task
	Post
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		JWTAuthorization: NewJWTAuthorizationService(repos.User),
		User:             NewUserService(repos.User),
		Task:             NewTaskService(repos.Task),
		Post:             NewPostService(repos.Post),
		Comment:          NewCommentService(repos.Comment),
	}
}
