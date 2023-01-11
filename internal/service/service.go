package service

import (
	"todo-list/internal/entity"
	"todo-list/internal/repository"
)

type JWTAuthorization interface {
	GenerateAccessRefreshTokens(id int) (string, string, error)
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
	CreateBulk(userId int, tasks []entity.TaskCreate) ([]int, error)
	GetById(id int) (entity.Task, error)
	GetByIds(ids []int) ([]entity.Task, error)
	Update(userId, taskId int, task entity.TaskUpdate) error
	Delete(userId, taskId int) error
	GetByUserId(id int) ([]entity.Task, error)
	ParseFile(path string) ([]entity.TaskCreate, error)
}

type Post interface {
	Create(userId int, task entity.PostCreate) (int, error)
	GetById(id int) (entity.Post, error)
	GetAll() ([]entity.Post, error)
	Update(userId, postId int, post entity.PostUpdate) error
	Delete(userId, postId int) error
	GetByUserId(id int) ([]entity.Post, error)
}

type Comment interface {
	Create(userId int, comment entity.CommentCreate) (int, error)
	GetById(id int) (entity.Comment, error)
	GetByPostId(id int) ([]entity.Comment, error)
	Update(userId, commentId int, comment entity.CommentUpdate) error
	Delete(userId, commentId int) error
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
		Post:             NewPostService(repos.Post, repos.Task),
		Comment:          NewCommentService(repos.Comment, repos.Post),
	}
}
