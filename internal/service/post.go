package service

import (
	"todo-list/internal/entity"
	"todo-list/internal/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (p *PostService) Create(userId int, post entity.PostCreate) (int, error) {
	return p.repo.Create(userId, post)
}

func (p *PostService) GetById(id int) (entity.Post, error) {
	return p.repo.GetById(id)
}

func (p *PostService) GetAll() ([]entity.Post, error) {
	return p.repo.GetAll()
}

func (p *PostService) Update(id int, post entity.PostUpdate) error {
	if err := post.Validate(); err != nil {
		return err
	}
	return p.repo.Update(id, post)
}

func (p *PostService) Delete(id int) error {
	return p.repo.Delete(id)
}

func (p *PostService) GetByUserId(id int) ([]entity.Post, error) {
	return p.repo.GetByUserId(id)
}
