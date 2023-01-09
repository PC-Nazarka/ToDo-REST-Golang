package service

import (
	"errors"
	"todo-list/internal/entity"
	"todo-list/internal/repository"
)

type PostService struct {
	repo     repository.Post
	taskRepo repository.Task
}

func NewPostService(repo repository.Post, taskRepo repository.Task) *PostService {
	return &PostService{repo: repo, taskRepo: taskRepo}
}

func (p *PostService) Create(userId int, post entity.PostCreate) (int, error) {
	task, err := p.taskRepo.GetById(post.TaskId)
	if err != nil {
		return 0, err
	}
	if userId != task.UserId {
		return 0, errors.New("you can't create post with task of another user")
	}
	return p.repo.Create(userId, post)
}

func (p *PostService) GetById(id int) (entity.Post, error) {
	return p.repo.GetById(id)
}

func (p *PostService) GetAll() ([]entity.Post, error) {
	return p.repo.GetAll()
}

func (p *PostService) Update(userId, postId int, post entity.PostUpdate) error {
	postExists, err := p.repo.GetById(postId)
	if err != nil {
		return err
	}
	if userId != postExists.UserId {
		return errors.New("you can't create post with task of another user")
	}
	if err = post.Validate(); err != nil {
		return err
	}
	return p.repo.Update(postId, post)
}

func (p *PostService) Delete(userId, postId int) error {
	post, err := p.repo.GetById(postId)
	if err != nil {
		return err
	}
	if userId != post.UserId {
		return errors.New("you can't create post with task of another user")
	}
	return p.repo.Delete(postId)
}

func (p *PostService) GetByUserId(id int) ([]entity.Post, error) {
	return p.repo.GetByUserId(id)
}
